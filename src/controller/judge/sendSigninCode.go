package judge_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	judge_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/judge"
	"github.com/matheuswww/quikworkout-games-backend/src/cookies/judge/judge_signin_cookie"
	judge_domain "github.com/matheuswww/quikworkout-games-backend/src/model/judge"
	"go.uber.org/zap"
)

func (jc *judgeController) SendSigninCode(c *gin.Context) {
	logger.Info("Init SendSigninCode Controller", zap.String("journey", "SendSigninCode Controller"))
	var judgeRequest judge_request.Signin
	if err := c.ShouldBindJSON(&judgeRequest); err != nil {
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	judgeDomain := judge_domain.NewJudgeSendSigninDomain(judgeRequest.Email, judgeRequest.Senha)
	restErr := jc.judgeService.SendSigninCode(judgeDomain, judgeRequest.Token)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	err := judge_signin_cookie.SendSigninCookie(c, judgeDomain.GetId())
	if err != nil {
		logger.Error("Error trying send cookie", err, zap.String("journey", "SendSigninCode Controller"))
		restErr := rest_err.NewInternalServerError("código gerado porém não foi possivel gerar sua sessão")
		c.JSON(restErr.Code, restErr)
		return
	}
	logger.Info("Signin code sended with success", zap.String("journey", "SendSigninCode Controller"))
	c.Status(http.StatusOK)
}
