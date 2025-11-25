package main

import (
	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/model"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&model.User{}, 
		&model.Court{}, 
		&model.Timeslot{}, 
		&model.Reservation{}, 
		&model.Payment{},
	)
}