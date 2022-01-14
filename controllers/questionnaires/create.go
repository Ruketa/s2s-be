package controllers

import (
	"goapi/db"
	model "goapi/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	//"fmt"

	"net/http"

	_ "github.com/lib/pq"
)

func CreateQuestionnaire(c *gin.Context) {
	do := db.GetDB()

	var q model.Questionnaire

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
