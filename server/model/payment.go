package model

import "gorm.io/gorm"

type PaymentStatus string
var (
	PaymentPending PaymentStatus = "pending"
	PaymentSettlement PaymentStatus = "settlement"
	PaymentCapture PaymentStatus = "capture"
	PaymentCancel PaymentStatus = "cancel"
	PaymentExpire PaymentStatus = "expire"
	PaymentDeny PaymentStatus = "deny"
)

type Payment struct {
	gorm.Model
	UserID        uint           `json:"user_id"`
	ReservationID uint           `json:"reservation_id"`
	Reservation   *Reservation   `json:"reservation" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount        float64        `json:"amount" gorm:"type:decimal(10,2)"`
	Status        PaymentStatus         `json:"status" gorm:"default:'pending'"`
	TransactionID string         `json:"transaction_id" gorm:"uniqueIndex"`
}