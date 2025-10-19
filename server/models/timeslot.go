package models

import "time"

type Timeslot struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	CreatedAt time.Time `json:"createdAt"`
}
