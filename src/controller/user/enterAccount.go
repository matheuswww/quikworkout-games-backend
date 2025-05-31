package user_controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	user_proflie_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_profile"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	model_util "github.com/matheuswww/quikworkout-games-backend/src/model/util"
	"go.uber.org/zap"
)

func (uc *userController) EnterAccount(c *gin.Context) {
	logger.Info("Init EnterAccount Controller")
	cookie, err := user_proflie_cookie.GetUserProfileCookieValues(c)
	if err != nil {
		logger.Error("Erro trying get cookie", err, zap.String("journey", "PayOrder Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}
	restErr := model_util.CheckUser(cookie.SessionId, cookie.Id)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	userDomain := user_domain.NewUserDomain(cookie.Id, "", "", "", 0, "", "")
	restErr = uc.userService.EnterAccount(userDomain)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	cookieErr := user_games_cookie.SendUserGamesCookie(c, userDomain.GetId(), userDomain.GetSessionId(), true)
	if cookieErr != nil {
		logger.Error("Error trying create session", cookieErr, zap.String("journey", "EnterAccount Controller"))
		restErr := rest_err.NewInternalServerError("não foi possível criar uma sessão")
		c.JSON(restErr.Code, restErr)
		return
	}
	c.Header("Access-Control-Expose-Headers", "Set-Cookie")
	logger.Info(fmt.Sprintf("EnterAccount success!,id: %s", userDomain.GetId()), zap.String("journey", "EnterAccount Controller"))
	c.Status(http.StatusOK)
}