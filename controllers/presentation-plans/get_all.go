package controllers

import (
	"encoding/json"
	"goapi/db"
	model "goapi/models"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func GetPresentationPlan(c *gin.Context) {
	do := db.GetDB()

	var ps []model.PresentationPlan

	do.Find(&ps)

	psJson, _ := json.Marshal(ps)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(psJson)
}
