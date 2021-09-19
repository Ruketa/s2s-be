package handlers

import (
	"goapi/db"
	"goapi/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

func PostStudySession(c *gin.Context) {
	do := db.GetDB()

	var ss entity.StudySession

	err := c.BindJSON(&ss)
	if err != nil {
		panic(err)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	ss.Id = id.String()

	do.Create(&ss)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
}
