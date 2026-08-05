package usecase

import (
	"context"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/model"
	"github.com/dandimuzaki/badminton-server/repository"
	"github.com/dandimuzaki/badminton-server/utils"
	"go.uber.org/zap"
)

type ReservationUsecase interface {
	CreateReservation(ctx context.Context, req dto.ReservationRequest) (*model.Reservation, error)
	GetReservationHistory(ctx context.Context) ([]model.Reservation, error)
	CancelReservation(ctx context.Context, id uint) error
}

type reservationUsecase struct {
	Repo *repository.Repository
	Log *zap.Logger
}

func NewReservationUsecase(repo *repository.Repository, log *zap.Logger) ReservationUsecase {
	return &reservationUsecase{
		Repo: repo,
		Log: log,
	}
}

func (u *reservationUsecase) CreateReservation(ctx context.Context, req dto.ReservationRequest) (*model.Reservation, error) {
	// Get user id
	userID, ok := ctx.Value("user_id").(uint)
	if !ok {
		return nil, utils.ErrInvalidUserID
	}
	
	// Check if timeslot already booked for that court
	existing, err := u.Repo.ReservationRepo.GetExistingReservation(ctx, req)
	if existing != nil {
		u.Log.Error("Court is already booked", zap.Error(err))
		return nil, utils.ErrCourtAlreadyBooked
	}

	// Create reservation with "pending" status
	reservation := model.Reservation{
		UserID:     userID,
		CourtID:    req.CourtID,
		Date:       req.Date,
		TimeSlotID: req.TimeSlotID,
		Status:     model.ReservationPending,
	}

	reservation, err = u.Repo.ReservationRepo.CreateReservation(ctx, &reservation)
	if err != nil {
		u.Log.Error("Failed to create reservation", zap.Error(err))
		return nil, err
	}

	return &reservation, nil
}

func (u *reservationUsecase) GetReservationHistory(ctx context.Context) ([]model.Reservation, error) {
	// Get user id
	userID, ok := ctx.Value("user_id").(uint)
	if !ok {
		return nil, utils.ErrInvalidUserID
	}

	reservations, err := u.Repo.ReservationRepo.GetReservationHistory(ctx, userID)
	if err != nil {
		u.Log.Error("Failed to get reservation history", zap.Error(err))
		return nil, err
	}

	return reservations, nil
}

func (u *reservationUsecase) CancelReservation(ctx context.Context, id uint) error {
	// Get reservation by id
	reservation, err := u.Repo.ReservationRepo.GetReservationByID(ctx, id)
	if err != nil {
		u.Log.Error("Failed to get reservation", zap.Error(err))
		return err
	}

	if reservation.Status != model.ReservationPending {
		return utils.ErrInvalidReservation
	}

	data := map[string]interface{}{
		"status": model.ReservationCancelled,
	}

	if err := u.Repo.ReservationRepo.UpdateReservation(ctx, id, data); err != nil {
		u.Log.Error("Failed to cancel reservation", zap.Error(err))
		return err
	}

	return nil
}