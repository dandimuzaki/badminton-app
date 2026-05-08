package model

import (
	"gorm.io/gorm"
)

type UserRole string
var (
	RoleCustomer UserRole = "customer"
)

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	Role UserRole `json:"role"`

	// Relations
	Reservations []Reservation `gorm:"foreignKey:UserID"`
}