package main

import "goapi/handlers"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main(){
	engine := gin.Default()

	// CORS
	engine.Use(cors.New(cors.Config{
		// アクセス許可をするドメイン
		AllowOrigins: []string{"http://localhost:3000"},
		// 許可するHTTPメソッド
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		// リクエスト結果のキャッシュ有効期限(秒)
		MaxAge: 12 * 60 * 60,
	}))

	group := engine.Group("/api")
	{
		group.GET("/questionnaire", handlers.Questionnaire)

		group.GET("/questionnaire/:holding_num", func(c *gin.Context) {
			holding_num := c.Param("holding_num")
			handlers.QuestionnaireByHoldingNum(c, holding_num)
		})
	}

	// start server
	engine.Run(":8000")
	fmt.Println("Server is running on port 8000")
}