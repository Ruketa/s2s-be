package controllers

import (
	"encoding/json"
	"goapi/db"
	model "goapi/models"

	"github.com/gin-gonic/gin"

	//"fmt"

	"net/http"

	_ "github.com/lib/pq"
)

func GetQuestionnaire(c *gin.Context) {
	do := db.GetDB()

	var qs []model.Questionnaire

	do.Find(&qs)

	qsJson, _ := json.Marshal(qs)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(qsJson)
}
