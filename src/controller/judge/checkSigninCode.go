package judge_controller

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	judge_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/judge"
	judge_profile_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/judge/judge_cookie"
	"github.com/matheuswww/quikworkout-games-backend/src/cookies/judge/judge_signin_cookie"
	judge_domain "github.com/matheuswww/quikworkout-games-backend/src/model/judge"
	"go.uber.org/zap"
)

func (jc *judgeController) CheckSigninCode(c *gin.Context) {
	logger.Info("Init CheckSigninCode Controller", zap.String("journey", "CheckSigninCode Controller"))
	cookie, err := judge_signin_cookie.GetSigninValues(c)
	if err != nil {
		logger.Error("Error trying get cookie", err, zap.String("journey", "CheckSigninCode Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}
	var adminAuthRequest judge_request.Code
	if err := c.ShouldBindJSON(&adminAuthRequest); err != nil {
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	adminAuthDomain := judge_domain.NewJudgeCheckSigninDomain(cookie.Id)
	restErr := jc.judgeService.CheckSigninCode(adminAuthDomain, adminAuthRequest.Codigo, adminAuthRequest.Token)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	err = judge_profile_cookie.SendJudgeCookie(c, cookie.Id)
	c.Header("Access-Control-Expose-Headers", "Set-Cookie")
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		logger.Error("Error trying load env", errors.New("error loading env"), zap.String("journey", "CheckSigninCode Controller"))
	}
	if err != nil {
		logger.Error("Error trying send cookie", err, zap.String("journey", "CheckSigninCode Controller"))
		restErr := rest_err.NewInternalServerError("código valido porém não foi possivel gerar sua sessão")
		c.SetCookie("adminAuthSignin", "", -1, "/judge", domain, true, true)
		c.JSON(restErr.Code, restErr)
		return
	}
	c.SetCookie("adminAuthSignin", "", -1, "/judge", domain, true, true)
	logger.Info("Signin code checked with success", zap.String("journey", "CheckSigninCode Controller"))
	c.Status(http.StatusOK)
}
