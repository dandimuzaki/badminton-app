package model

import (
	"gorm.io/gorm"
)

type Timeslot struct {
	gorm.Model
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
