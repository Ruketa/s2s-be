package model

import (
	"time"
)

// Questionnaire
type Questionnaire struct {
	Id                   string    `json:"id"`
	Satisfaction_level   int       `json:"satisfaction_level" validate:"required, gte=0, lte=5"`
	Recommendation_level int       `json:"recommendation_level" validate:"required, gte=0, lte=5"`
	Topics               string    `json:"topics" validate:"required"`
	Participation_level  int       `json:"participation_level" validate:"required, gte=0, lte=5"`
	Presentation_level   int       `json:"presentation_level" validate:"required, gte=0, lte=5"`
	Free_comment         string    `json:"free_comment" validate:"required"`
	Holding_num          int       `json:"holding_num" validate:"required, gte=0"`
	Created_at           time.Time `json:"created_at"`
}
