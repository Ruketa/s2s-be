package entity

import (
	"time"
)

// Questionnaire entity
type Questionnaire struct {
	Id                   string    `json:"id"`
	Satisfaction_level   int       `json:"satisfaction_level" binding:"required, min=0, max=5"`
	Recommendation_level int       `json:"recommendation_level" binding:"required, min=0, max=5"`
	Topics               string    `json:"topics" binding:"required"`
	Participation_level  int       `json:"participation_level" binding:"required, min=0, max=5"`
	Presentation_level   int       `json:"presentation_level" binding:"required, min=0, max=5"`
	Free_comment         string    `json:"free_comment" binding:"required"`
	Holding_num          int       `json:"holding_num" binding:"required, min=0"`
	Created_at           time.Time `json:"created_at"`
}

// PresentationPan entity
type PresentationPlan struct {
	Id                 string    `json:"id"`
	Presenter          string    `json:"presenter" binding:"required"`
	Presentation_title string    `json:"presentation_title" binding:"required"`
	Scheduled_on       time.Time `json:"scheduled_on" binding:"required"`
	Division           string    `json:"division" binding:"required"`
	Created_at         time.Time `json:"created_at"`
}

// StudySession entity
type StudySession struct {
	Id          string    `json:"id"`
	Holding_num int       `json:"holding_num" binding:"required"`
	Plan_id     string    `json:"plan_id" binding:"required"`
	Evented_on  time.Time `json:"evented_on" binding:"required"`
	Created_at  time.Time `json:"created_at"`
}
