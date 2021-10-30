package handlers

import (
	"encoding/json"
	"goapi/db"
	"goapi/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	//"fmt"

	"net/http"

	_ "github.com/lib/pq"
)

func Questionnaire(c *gin.Context) {
	do := db.GetDB()

	var qs []entity.Questionnaire

	do.Find(&qs)

	qsJson, _ := json.Marshal(qs)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(qsJson)
}

func QuestionnaireByHoldingNum(c *gin.Context) {
	do := db.GetDB()

	var qs []entity.Questionnaire

	do.Where("holding_num = ?", c.Param("holding_num")).Find(&qs)

	qsJson, _ := json.Marshal(qs)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(qsJson)
}

func PostQuestionnaire(c *gin.Context) {
	do := db.GetDB()

	var q entity.Questionnaire

	err := c.BindJSON(&q)
	if err != nil {
		panic(err)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	q.Id = id.String()

	do.Create(&q)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
}
