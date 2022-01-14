package controllers

import (
	"encoding/json"
	"goapi/db"
	model "goapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQuestionnaireByHoldingNum(c *gin.Context) {
	do := db.GetDB()

	var qs []model.Questionnaire

	do.Where("holding_num = ?", c.Param("holding_num")).Find(&qs)

	qsJson, _ := json.Marshal(qs)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(qsJson)
}
