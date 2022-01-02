package server

import (
	"goapi/router"

	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()

	router.Setup(engine)

	// start server
	engine.Run(":8000")
}
