package adaptor

import (
	"github.com/dandimuzaki/badminton-server/usecase"
	"go.uber.org/zap"
)

type Handler struct{
	AuthHandler AuthHandler
	CourtHandler CourtHandler
	TimeslotHandler TimeslotHandler
	ReservationHandler ReservationHandler
	PaymentHandler PaymentHandler
}

func NewHandler(u *usecase.Usecase, log *zap.Logger) Handler {
	return Handler{
		AuthHandler: NewAuthHandler(u.AuthUsecase, log),
		CourtHandler: NewCourtHandler(u.CourtUsecase, log),
		TimeslotHandler: NewTimeslotHandler(u.TimeslotUsecase, log),
		ReservationHandler: NewReservationHandler(u.ReservationUsecase, log),
		PaymentHandler: NewPaymentHandler(u.PaymentUsecase, log),
	}
}