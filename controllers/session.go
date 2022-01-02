package controllers

import (
	"encoding/json"
	"fmt"
	"goapi/db"
	"goapi/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

func CreateStudySession(c *gin.Context) {
	do := db.GetDB()

	var ss entity.StudySession

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

func GetStudySession(c *gin.Context) {
	do := db.GetDB()

	var ss []entity.StudySession

	do.Find(&ss)

	ssJson, err := json.Marshal(ss)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(ssJson))

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.JSON(http.StatusOK, ssJson)
}
