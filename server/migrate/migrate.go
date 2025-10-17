package main

import (
	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.User{}, 
		&models.Court{}, 
		&models.Timeslot{}, 
		&models.Reservation{}, 
		&models.Payment{},
	)
}