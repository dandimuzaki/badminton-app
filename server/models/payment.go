package models

import "time"

type Payment struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	ReservationID  uint      `json:"reservationId"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	TransactionID  string    `json:"transactionId"`
	CreatedAt      time.Time `json:"createdAt"`
}
