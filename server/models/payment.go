package models

import "time"

type Payment struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	UserID        uint           `json:"userId"`
	ReservationID uint           `json:"reservationId"`
	Reservation   *Reservation   `json:"reservation" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount        float64        `json:"amount" gorm:"type:decimal(10,2)"`
	Status        string         `json:"status" gorm:"default:'pending'"`
	TransactionID string         `json:"transactionId" gorm:"uniqueIndex"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}



