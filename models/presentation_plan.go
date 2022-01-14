package model

import (
	"time"
)

// PresentationPan entity
type PresentationPlan struct {
	Id                 string    `json:"id"`
	Presenter          string    `json:"presenter" binding:"required"`
	Presentation_title string    `json:"presentation_title" binding:"required"`
	Scheduled_on       time.Time `json:"scheduled_on" binding:"required"`
	Division           string    `json:"division" binding:"required"`
	Created_at         time.Time `json:"created_at"`
}
