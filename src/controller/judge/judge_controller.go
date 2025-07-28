package judge_controller

import (
	"github.com/gin-gonic/gin"
	judge_service "github.com/matheuswww/quikworkout-games-backend/src/model/judge/service"
)

func NewJudgeController(judgeService judge_service.JudgeService) JudgeController {
	return &judgeController{
		judgeService,
	}
}

type judgeController struct {
	judgeService judge_service.JudgeService
}

type JudgeController interface {
	CheckSigninCode(c *gin.Context)
	SendSigninCode(c *gin.Context)
}