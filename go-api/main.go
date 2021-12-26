package main

import (
	"fmt"
	"goapi/db"
	"goapi/handlers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	engine := gin.Default()

	// CORS
	engine.Use(cors.New(cors.Config{
		// アクセス許可をするドメイン
		AllowOrigins: []string{"*"},
		// 許可するHTTPメソッド
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		// リクエスト結果のキャッシュ有効期限(秒)
		MaxAge: 12 * 60 * 60,
	}))

	group := engine.Group("/api")
	{
		group.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "ok"})
		})
		group.GET("/questionnaire", handlers.Questionnaire)
		group.GET("/questionnaire/:holding_num", handlers.QuestionnaireByHoldingNum)
		group.POST("/questionnaire", handlers.PostQuestionnaire)

		group.GET("/presenter", handlers.GetPresenter)
		group.GET("/presenter/:holding_num", handlers.GetPresenterByHoldingNum)
		group.POST("/presentation_plan", handlers.PostPresentationPlan)

		group.POST("/session", handlers.PostStudySession)
	}

	return engine
}

func main() {
	db.InitDB()
	defer db.CloseDB()

	// start server
	router().Run(":8000")
	fmt.Println("Server is running on port 8000")
}
