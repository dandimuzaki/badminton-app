package repository

import (
	"context"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/infra"
	"github.com/dandimuzaki/badminton-server/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ReservationRepository interface {
	CreateReservation(ctx context.Context, reservation *model.Reservation) error
	GetReservationHistory(ctx context.Context, userID uint) ([]model.Reservation, error)
	GetReservationByID(ctx context.Context, id uint) (*model.Reservation, error)
	UpdateReservation(ctx context.Context, id uint, data map[string]interface{}) error
	UpdateReservationByTransactionID(ctx context.Context, transactionID string, status string) error
	GetExistingReservation(ctx context.Context, req dto.ReservationRequest) (*model.Reservation, error)
}

type reservationRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewReservationRepository(db *gorm.DB, log *zap.Logger) ReservationRepository {
	return &reservationRepository{
		DB:  db,
		Log: log,
	}
}

func (r *reservationRepository) CreateReservation(ctx context.Context, reservation *model.Reservation) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.Create(reservation).Error; err != nil {
		r.Log.Error("Failed to create reservation", zap.Error(err))
		return err
	}
	return nil
}

func (r *reservationRepository) GetReservationHistory(ctx context.Context, userID uint) ([]model.Reservation, error) {
	db := infra.GetDB(ctx, r.DB)
	var reservations []model.Reservation
	if err := db.Model(&model.Reservation{}).
		Where("user_id = ?", userID).
		Preload("User").
		Preload("Court").
		Preload("Timeslot").
		Preload("Payment").
		Order("date DESC").
		Find(&reservations).Error; err != nil {
		r.Log.Error("Failed to get reservation history", zap.Error(err))
		return nil, err
	}
	return reservations, nil
}

func (r *reservationRepository) GetReservationByID(ctx context.Context, id uint) (*model.Reservation, error) {
	db := infra.GetDB(ctx, r.DB)
	var reservation model.Reservation
	if err := db.
		Preload("User").
		Preload("Court").
		Preload("Timeslot").
		Preload("Payment").
		First(&reservation, id).Error; err != nil {
		r.Log.Error("Failed to get reservation by ID", zap.Error(err))
		return nil, err
	}
	return &reservation, nil
}

func (r *reservationRepository) UpdateReservation(ctx context.Context, id uint, data map[string]interface{}) (error) {
	db := infra.GetDB(ctx, r.DB)
	if err := db.Model(&model.Reservation{}).Where("id = ?", id).Updates(data).Error; err != nil {
		r.Log.Error("Failed to update reservation", zap.Error(err))
		return err
	}
	return nil
}

func (r *reservationRepository) UpdateReservationByTransactionID(ctx context.Context, transactionID string, status string) error {
	db := infra.GetDB(ctx, r.DB)
	if err := db.Exec(`
			UPDATE reservations 
			SET status = ?
			WHERE id = (
				SELECT reservation_id FROM payments WHERE transaction_id = ?
			)
		`, status, transactionID).Error; err != nil {
		r.Log.Error("Failed to update reservation", zap.Error(err))
		return err
	}

	return nil
}

func (r *reservationRepository) GetExistingReservation(ctx context.Context, req dto.ReservationRequest) (*model.Reservation, error) {
	db := infra.GetDB(ctx, r.DB)
	// Check if timeslot already booked for that court
	var existing model.Reservation
	if err := db.
		Where("court_id = ? AND date = ? AND time_slot_id = ?", req.CourtID, req.Date, req.TimeSlotID).
		First(&existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}