package handlers

import (
	"encoding/json"
	"goapi/db"
	"goapi/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type PresentationResult struct {
	Holding_num        string
	Event_date         time.Time
	Presenter          string
	Presentation_title string
	Division           string
}

func GetPresenter(c *gin.Context) {
	do := db.GetDB()

	ps := []entity.PresentationPlan{}

	do.Find(&ps)

	psJson, _ := json.Marshal(ps)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(psJson)
}

func GetPresenterByHoldingNum(c *gin.Context) {
	do := db.GetDB()

	results := []PresentationResult{}
	do.Table("study_session").
		Select("study_session.holding_num, study_session.event_date, presentation_plan.presenter, presentation_plan.presentation_title, presentation_plan.division").
		Joins("left join presentation_plan on study_session.presenter_id = presentation_plan.id").
		Where("study_session.holding_num = ?", c.Param("holding_num")).
		Scan(&results)

	resultsJson, _ := json.Marshal(results)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(resultsJson)
}

func PostPresentationPlan(c *gin.Context) {
	db := db.GetDB()

	var p entity.PresentationPlan

	err := c.BindJSON(&p)
	if err != nil {
		panic(err)
	}

	// UUID
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	p.Id = id.String()

	// Created at
	p.Created_at = time.Now()

	db.Create(&p)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
}
