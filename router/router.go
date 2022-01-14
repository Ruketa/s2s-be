package router

import (
	"fmt"

	presentation_plan "goapi/controllers/presentation-plans"
	questionnaire "goapi/controllers/questionnaires"
	session "goapi/controllers/sessions"

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
	group_q := group.Group("/questionnaire")
	{
		group_q.GET("", questionnaire.GetQuestionnaire)
		group_q.GET("/:holding_num", questionnaire.GetQuestionnaireByHoldingNum)
		group_q.POST("", questionnaire.CreateQuestionnaire)
	}
	group_p := group.Group("/presentation_plan")
	{
		group_p.GET("", presentation_plan.GetPresentationPlan)
		group_p.POST("", presentation_plan.CreatePresentationPlan)
	}
	group_s := group.Group("/study_session")
	{
		group_s.GET("", session.GetStudySession)
		group_s.POST("", session.CreateStudySession)
	}

	engine.GET("/healthcheck", func(c *gin.Context) {
		fmt.Println("healthcheck")
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
