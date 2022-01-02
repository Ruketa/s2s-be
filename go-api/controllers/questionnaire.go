package controllers

import (
	"encoding/json"
	"goapi/db"
	"goapi/entity"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	//"fmt"

	"net/http"

	_ "github.com/lib/pq"
)

func GetQuestionnaire(c *gin.Context) {
	do := db.GetDB()

	var qs []entity.Questionnaire

	do.Find(&qs)

	qsJson, _ := json.Marshal(qs)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(qsJson)
}

func GetQuestionnaireByHoldingNum(c *gin.Context) {
	do := db.GetDB()

	var qs []entity.Questionnaire

	do.Where("holding_num = ?", c.Param("holding_num")).Find(&qs)

	qsJson, _ := json.Marshal(qs)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(qsJson)
}

func CreateQuestionnaire(c *gin.Context) {
	do := db.GetDB()

	var q entity.Questionnaire

	err := c.ShouldBindJSON(&q)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate UUID
	id, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		panic(err)
	}
	q.Id = id.String()

	// Created at
	q.Created_at = time.Now()

	do.Create(&q)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
}
