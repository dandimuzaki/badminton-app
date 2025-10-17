package models

import "time"

type Court struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Price 		float64 	`json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}
