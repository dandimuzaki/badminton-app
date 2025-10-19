package models

import "time"

type Reservation struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	UserID     uint       `json:"userId"`
	User       User       `gorm:"foreignKey:UserID"`
	CourtID    uint       `json:"courtId"`
	Court      Court      `gorm:"foreignKey:CourtID"`
	Date       time.Time  `json:"date"`
	TimeSlotID uint       `json:"timeSlotId"`
	Timeslot   Timeslot   `gorm:"foreignKey:TimeSlotID"`
	Status     string     `json:"status"` // pending, confirmed, canceled, completed
	Payment    *Payment   `gorm:"foreignKey:ReservationID"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}

