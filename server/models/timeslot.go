package models

import "time"

type Timeslot struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	CreatedAt time.Time `json:"createdAt"`
}
