package dto

import (
	"time"

	"github.com/dandimuzaki/badminton-server/model"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AvailableCourtQuery struct {
	Date       string `form:"date"`
	TimeSlotID uint   `form:"time_slot_id"`
}

type AvailableCourtRequest struct {
	Date       time.Time
	TimeSlotID uint
}

func ToCourtRequest(query *AvailableCourtQuery) (*AvailableCourtRequest, error) {
	var date time.Time
	if query.Date != "" {
		time, err := time.Parse("2006-01-02", query.Date)
		if err != nil {
			return nil, err
		}
		date = time
	}
	
	return &AvailableCourtRequest{
		Date: date,
		TimeSlotID: query.TimeSlotID,
	}, nil
}

type AvailableTimeslotQuery struct {
	Date       string `form:"date"`
}

type AvailableTimeslotRequest struct {
	Date       time.Time
}

func ToTimeslotRequest(query *AvailableTimeslotQuery) (*AvailableTimeslotRequest, error) {
	var date time.Time
	if query.Date != "" {
		time, err := time.Parse("2006-01-02", query.Date)
		if err != nil {
			return nil, err
		}
		date = time
	}
	
	return &AvailableTimeslotRequest{
		Date: date,
	}, nil
}

type ReservationQuery struct {
	CourtID    uint   `json:"court_id" binding:"required"`
	Date       string `json:"date" binding:"required"`
	TimeSlotID uint   `json:"time_slot_id" binding:"required"`
}

type ReservationRequest struct {
	CourtID    uint   `json:"court_id" binding:"required"`
	Date       time.Time `json:"date" binding:"required"`
	TimeSlotID uint   `json:"time_slot_id" binding:"required"`
}

func ToReservationRequest(query *ReservationQuery) (*ReservationRequest, error) {
	var date time.Time
	if query.Date != "" {
		time, err := time.Parse("2006-01-02", query.Date)
		if err != nil {
			return nil, err
		}
		date = time
	}
	
	return &ReservationRequest{
		Date: date,
		CourtID: query.CourtID,
		TimeSlotID: query.TimeSlotID,
	}, nil
}

type PaymentRequest struct {
	ReservationID uint `json:"reservation_id"`
	Amount float64 `json:"amount"`
}

type PaymentCallback struct {
	OrderID string `json:"order_id"`
	Status model.PaymentStatus `json:"transaction_status"`
}