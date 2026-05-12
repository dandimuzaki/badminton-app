package repository

import (
	"context"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/infra"
	"github.com/dandimuzaki/badminton-server/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CourtRepository interface {
	GetAvailableCourts(ctx context.Context, req dto.AvailableCourtRequest) ([]model.Court, error)
	GetAllCourts(ctx context.Context) ([]model.Court, error)
}

type courtRepository struct {
	DB *gorm.DB
	Log *zap.Logger
}

func NewCourtRepository(db *gorm.DB, log *zap.Logger) CourtRepository {
	return &courtRepository{
		DB: db,
		Log: log,
	}
}

func (r *courtRepository) GetAvailableCourts(ctx context.Context, req dto.AvailableCourtRequest) ([]model.Court, error) {
	db := infra.GetDB(ctx, r.DB)
	var courts []model.Court

	// Use subquery to exclude booked courts
	subQuery := db.Model(&model.Reservation{}).
		Select("court_id").
		Where("date = ? AND time_slot_id = ? AND status = ?", req.Date, req.TimeSlotID, model.ReservationPending)

	if err := db.
		Where("id NOT IN (?)", subQuery).
		Find(&courts).Error; err != nil {
		return nil, err
	}

	return courts, nil
}

func (r *courtRepository) GetAllCourts(ctx context.Context) ([]model.Court, error) {
	db := infra.GetDB(ctx, r.DB)
	var courts []model.Court

	if err := db.
		Find(&courts).Error; err != nil {
		return nil, err
	}

	return courts, nil
}