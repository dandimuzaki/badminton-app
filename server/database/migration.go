package database

import (
	"github.com/dandimuzaki/badminton-server/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{}, 
		&model.Court{}, 
		&model.Timeslot{}, 
		&model.Reservation{}, 
		&model.Payment{},
	)
}
