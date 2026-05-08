package usecase

import (
	"github.com/dandimuzaki/badminton-server/repository"
	"github.com/dandimuzaki/badminton-server/utils"
	"go.uber.org/zap"
)

type Usecase struct{
	AuthUsecase AuthUsecase
	CourtUsecase CourtUsecase
	TimeslotUsecase TimeslotUsecase
	ReservationUsecase ReservationUsecase
	PaymentUsecase PaymentUsecase
}

func NewUsecase(tx TxManager, repo *repository.Repository, tokenService TokenService, log *zap.Logger, config utils.Configuration) *Usecase {
	return &Usecase{
		AuthUsecase: NewAuthUsecase(repo, tokenService, log),
		CourtUsecase: NewCourtUsecase(repo, log),
		TimeslotUsecase: NewTimeslotUsecase(repo, log),
		ReservationUsecase: NewReservationUsecase(repo, log),
		PaymentUsecase: NewPaymentUsecase(repo, log, config.Midtrans),
	}
}