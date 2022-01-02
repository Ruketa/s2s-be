package controllers

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

func GetPresentationPlan(c *gin.Context) {
	do := db.GetDB()

	var ps []entity.PresentationPlan

	do.Find(&ps)

	psJson, _ := json.Marshal(ps)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(psJson)
}

func CreatePresentationPlan(c *gin.Context) {
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
