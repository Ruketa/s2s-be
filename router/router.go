package router

import (
	"fmt"
	"goapi/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setting(engine *gin.Engine) {

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

}

func Setup(engine *gin.Engine) {

	setting(engine)

	group := engine.Group("/api")
	questionnaire := group.Group("/questionnaire")
	{
		questionnaire.GET("", controllers.GetQuestionnaire)
		questionnaire.GET("/:holding_num", controllers.GetQuestionnaireByHoldingNum)
		questionnaire.POST("", controllers.CreateQuestionnaire)
	}
	presentationPlan := group.Group("/presentation_plan")
	{
		presentationPlan.GET("", controllers.GetPresentationPlan)
		presentationPlan.POST("", controllers.CreatePresentationPlan)

	}
	studySession := group.Group("/study_session")
	{
		studySession.GET("", controllers.GetStudySession)
		studySession.POST("", controllers.CreateStudySession)
	}

	engine.GET("/healthcheck", func(c *gin.Context) {
		fmt.Println("healthcheck")
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
