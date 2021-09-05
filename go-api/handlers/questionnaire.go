package handlers

import "goapi/db"

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"encoding/json"
	"net/http"
)

func Questionnaire(c *gin.Context) {
	do := db.Connect()
	qs := db.SelectQuesionnaire(do)
	db.Close(do)

	qsJson, _ := json.Marshal(qs)
	
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(qsJson)
}

func QuestionnaireByHoldingNum(c *gin.Context, holding_num string) {
	do := db.Connect()
	qs := db.SelectQuesionnaireByHoldingNum(do, holding_num)
	db.Close(do)

	qsJson, _ := json.Marshal(qs)
	
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(qsJson)
}

