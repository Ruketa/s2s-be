package entity

import (
	"time"
)

// Questionnaire entity
type Questionnaire struct {
	Id                   string    `json:"id"`
	Satisfaction_level   int       `json:"satisfaction_level"`
	Recommendation_level int       `json:"recommendation_level"`
	Topics               string    `json:"topics"`
	Participation_level  int       `json:"participation_level"`
	Presentation_level   int       `json:"presentation_level"`
	Free_comment         string    `json:"free_comment"`
	Holding_num          int       `json:"holding_num"`
	Created_On           time.Time `json:"created_on"`
}

// PresentationPan entity
type PresentationPlan struct {
	Id                 string    `json:"id"`
	Presenter          string    `json:"presenter"`
	Presentation_title string    `json:"presentation_title"`
	Scheduled_on       time.Time `json:"scheduled_on"`
	Created_on         time.Time `json:"created_on"`
	Division           string    `json:"division"`
}

// StudySession entity
type StudySession struct {
	Id           string    `json:"id"`
	Holding_num  int       `json:"holding_num"`
	Presenter_id string    `json:"presenter_id"`
	Created_On   time.Time `json:"created_on"`
}
