package repository

import (
	"context"

	"github.com/dandimuzaki/badminton-server/dto"
	"github.com/dandimuzaki/badminton-server/infra"
	"github.com/dandimuzaki/badminton-server/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TimeslotRepository interface {
	GetAvailableTimeslots(ctx context.Context, req dto.AvailableTimeslotRequest) ([]model.Timeslot, error)
}

type timeslotRepository struct {
	DB *gorm.DB
	Log *zap.Logger
}

func NewTimeslotRepository(db *gorm.DB, log *zap.Logger) TimeslotRepository {
	return &timeslotRepository{
		DB: db,
		Log: log,
	}
}

func (r *timeslotRepository) GetAvailableTimeslots(ctx context.Context, req dto.AvailableTimeslotRequest) ([]model.Timeslot, error) {
	db := infra.GetDB(ctx, r.DB)
	var timeslots []model.Timeslot

	// First, count total courts
	var totalCourts int64
	db.Model(&model.Court{}).Count(&totalCourts)

	// Subquery: timeslot IDs that are fully booked
	subQuery := db.
		Model(&model.Reservation{}).
		Select("time_slot_id").
		Where("date = ? AND status = ?", req.Date, model.ReservationPending).
		Group("time_slot_id").
		Having("COUNT(court_id) >= ?", totalCourts)

	// Fetch timeslots that are not fully booked
	if err := db.
		Where("id NOT IN (?)", subQuery).
		Find(&timeslots).Error; err != nil {
		return nil, err
	}

	return timeslots, nil
}