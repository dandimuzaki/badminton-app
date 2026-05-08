package model

import (
	"time"

	"gorm.io/gorm"
)

type ReservationStatus string
var (
	ReservationPending ReservationStatus = "pending"
	ReservationPaid ReservationStatus = "paid"
	ReservationCheckedIn ReservationStatus = "checked in"
	ReservationCancelled ReservationStatus = "cancelled"
	ReservationCompleted ReservationStatus = "completed"
	ReservationFailed ReservationStatus = "failed"
)

type Reservation struct {
	gorm.Model
	UserID     uint       `json:"user_id"`
	CourtID    uint       `json:"court_id"`
	Date       time.Time  `json:"date"`
	TimeSlotID uint       `json:"time_slot_id"`
	Status     ReservationStatus     `json:"status"`

	// Relations
	User       User       `gorm:"foreignKey:UserID" json:"user"`
	Court      Court      `gorm:"foreignKey:CourtID" json:"court"`
	Timeslot   Timeslot   `gorm:"foreignKey:TimeSlotID" json:"timeslot"`
	Payment    *Payment   `gorm:"foreignKey:ReservationID" json:"payment"`
}

