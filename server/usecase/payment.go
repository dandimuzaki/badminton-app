package usecase

import (
	"context"
	"fmt"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/model"
	"github.com/dandimuzaki/badminton-server/repository"
	"github.com/dandimuzaki/badminton-server/utils"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"go.uber.org/zap"
)

type PaymentUsecase interface {
	CreatePayment(ctx context.Context, req dto.PaymentRequest) (*dto.CreatePaymentResponse, error)
	PaymentCallback(ctx context.Context, req dto.PaymentCallback) error
}

type paymentUsecase struct {
	Repo *repository.Repository
	Log  *zap.Logger
	Config utils.MidtransConfig
}

func NewPaymentUsecase(repo *repository.Repository, log *zap.Logger, config utils.MidtransConfig) PaymentUsecase {
	return &paymentUsecase{
		Repo: repo,
		Log:  log,
		Config: config,
	}
}

func (u *paymentUsecase) CreatePayment(ctx context.Context, req dto.PaymentRequest) (*dto.CreatePaymentResponse, error) {
	// Get reservation details
	reservation, err := u.Repo.ReservationRepo.GetReservationByID(ctx, req.ReservationID)
	if err != nil {
		u.Log.Error("Failed to get reservation details", zap.Error(err))
		return nil, err
	}

	// Get user id
	userID, ok := ctx.Value("user_id").(uint)
	if !ok {
		return nil, utils.ErrInvalidUserID
	}

	// Get user information
	user, err := u.Repo.UserRepo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, utils.ErrUserNotFound
	}

	// Validate reservation owner
	if reservation.UserID != userID {
		return nil, utils.ErrUnauthorized
	}

	// Initialize Midtrans client
	var s = snap.Client{}
	s.New(u.Config.ServerKey, midtrans.Sandbox)

	orderID := fmt.Sprintf("ORDER-%d-%d", reservation.ID, userID)

	// Construct transaction request to midtrans
	transaction := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: orderID,
			GrossAmt: int64(req.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    fmt.Sprintf("court-%d", reservation.CourtID),
				Price: int64(reservation.Court.Price),
				Qty:   1,
				Name:  reservation.Court.Name,
			},
		},
	}

	// Create transaction
	snapResp, err := s.CreateTransaction(transaction)

	// Save payment record into database
	payment := model.Payment{
		UserID: userID,
		ReservationID: req.ReservationID,
		Amount: req.Amount,
		Status: model.PaymentPending,
		TransactionID: orderID,
	}

	err = u.Repo.PaymentRepo.CreatePayment(ctx, &payment)
	if err != nil {
		u.Log.Error("Failed to save payment record into database", zap.Error(err))
		return nil, err
	}

	return &dto.CreatePaymentResponse{
		SnapToken: snapResp.Token,
		RedirectURL: snapResp.RedirectURL,
	}, nil
}

func (u *paymentUsecase) PaymentCallback(ctx context.Context, req dto.PaymentCallback) error {
	// Update payment record
	err := u.Repo.PaymentRepo.PaymentCallback(ctx, req)
	if err != nil {
		u.Log.Error("Failed to update payment record", zap.Error(err))
		return err
	}

	// Update reservation status
	if req.Status == model.PaymentSettlement || req.Status == model.PaymentCapture {
		err = u.Repo.ReservationRepo.UpdateReservationByTransactionID(ctx, req.OrderID, string(model.ReservationPaid))
	} else if req.Status == model.PaymentCancel || req.Status == model.PaymentExpire || req.Status == model.PaymentDeny {
		err = u.Repo.ReservationRepo.UpdateReservationByTransactionID(ctx, req.OrderID, string(model.ReservationFailed))
	}
	if err != nil {
		u.Log.Error("Failed to update reservation record", zap.Error(err))
		return err
	}

	return nil
}