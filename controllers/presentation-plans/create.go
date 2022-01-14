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

func CreatePresentationPlan(c *gin.Context) {
	db := db.GetDB()

	var p model.PresentationPlan

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
