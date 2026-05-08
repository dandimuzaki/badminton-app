package model

import (
	"gorm.io/gorm"
)

type Court struct {
	gorm.Model
	Name      string    `json:"name"`
	ImageURL	string 		`json:"image_url"`
	Type      string    `json:"type"`
	Description      string    `json:"description"`
	Location  string    `json:"location"`
	Price 		float64 	`json:"price"`
}
