package model

import (
	"time"
)

// StudySession entity
type StudySession struct {
	Id          string    `json:"id"`
	Holding_num int       `json:"holding_num" binding:"required"`
	Plan_id     string    `json:"plan_id" binding:"required"`
	Evented_on  time.Time `json:"evented_on" binding:"required"`
	Created_at  time.Time `json:"created_at"`
}
