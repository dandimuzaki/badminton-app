package models

import "time"

type Reservation struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"userId"`
	CourtID     uint      `json:"courtId"`
	Date       time.Time `json:"date" gorm:"type:date"`
	TimeSlotID  uint      `json:"timeSlotId"`
	Status      string    `json:"status"`
	Payment     Payment   `json:"payment" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentID   *uint     `json:"paymentId"`
	CreatedAt   time.Time `json:"createdAt"`
}
