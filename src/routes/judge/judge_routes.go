package judge_router

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	judge_controller "github.com/matheuswww/quikworkout-games-backend/src/controller/judge"
	judge_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/judge"
	judge_profile_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/judge/judge_cookie"
	"github.com/matheuswww/quikworkout-games-backend/src/cookies/judge/judge_signin_cookie"
	judge_repository "github.com/matheuswww/quikworkout-games-backend/src/model/judge/repository"
	judge_service "github.com/matheuswww/quikworkout-games-backend/src/model/judge/service"
	admin_router "github.com/matheuswww/quikworkout-games-backend/src/routes/admin"
	"go.uber.org/zap"
)

func InitJudgeRoutes(r *gin.RouterGroup, database *sql.DB) {
	judgeController := initJudgeRoutes(database)
	judgeGroup := r.Group("/judge")
	cookieStore, err := judge_cookie.Store()
	if err != nil {
		logger.Error("Error loading cookie store", err, zap.String("journey", "InitJudgeRoutes"))
		log.Fatal("Error cookie store")
	}
	sessionNames := []string{judge_signin_cookie.SessionSignin, judge_profile_cookie.SessionJudge}
	judgeGroup.Use(sessions.SessionsMany(sessionNames, cookieStore))

	judgeGroup.POST("/auth/sendSigninCode", judgeController.SendSigninCode)
	judgeGroup.POST("/auth/checkSigninCode", judgeController.CheckSigninCode)
	
	adminController := admin_router.GetAdminController(database)
	judgeGroup.GET("/getParticipants", authMiddleware, adminController.GetParticipants)
	judgeGroup.POST("/checkVideo", authMiddleware, adminController.CheckVideo)
	judgeGroup.POST("/desqualifyVideo", authMiddleware, adminController.DesqualifyVideo)
	judgeGroup.POST("/putTime", authMiddleware, adminController.PutTime)
	judgeGroup.POST("/putNoreps", authMiddleware, adminController.PutNoReps)
}

func authMiddleware(c *gin.Context) {
	_, err := judge_profile_cookie.GetJudgeCookieValues(c)
	if err != nil {
		logger.Error("Error trying get cookie", err, zap.String("journey", "judge route"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		c.Abort()
		return
	}
	c.Next()
}

func initJudgeRoutes(database *sql.DB) judge_controller.JudgeController {
	judgeRepository := judge_repository.NewJudgeRepository(database)
	judgeService := judge_service.NewJudgeService(judgeRepository)
	judgeController := judge_controller.NewJudgeController(judgeService)
	return judgeController
}
