package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo UserRepository
	CourtRepo CourtRepository
	TimeslotRepo TimeslotRepository
	ReservationRepo ReservationRepository
	PaymentRepo PaymentRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) *Repository {
	return &Repository{
		UserRepo: NewUserRepository(db, log),
		CourtRepo: NewCourtRepository(db, log),
		TimeslotRepo: NewTimeslotRepository(db, log),
		ReservationRepo: NewReservationRepository(db, log),
		PaymentRepo: NewPaymentRepository(db, log),
	}
}