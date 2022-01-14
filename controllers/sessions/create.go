package controllers

import (
	"goapi/db"
	model "goapi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

func CreateStudySession(c *gin.Context) {
	do := db.GetDB()

	var ss model.StudySession

	err := c.BindJSON(&ss)
	if err != nil {
		panic(err)
	}

	// Generate UUID
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	ss.Id = id.String()

	// Created at
	ss.Created_at = time.Now()

	do.Create(&ss)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
}
