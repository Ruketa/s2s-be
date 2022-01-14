package controllers

import (
	"encoding/json"
	"fmt"
	"goapi/db"
	model "goapi/models"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func GetStudySession(c *gin.Context) {
	do := db.GetDB()

	var ss []model.StudySession

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
