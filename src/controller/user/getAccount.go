package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/response"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	user_service_util "github.com/matheuswww/quikworkout-games-backend/src/model/user/service/util"
	"go.uber.org/zap"
)

func (uc *userController) GetAccount(c *gin.Context) {
	logger.Info("Init GetAccount Controller", zap.String("journey", "GetAccount Controller"))
	cookie, err := user_games_cookie.GetUserGamesCookieValues(c)
	if err != nil {
		logger.Error("Erro trying get cookie", err, zap.String("journey", "GetAccount Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}
	userDomain := user_domain.NewUserDomain(cookie.Id, "", "", "", 0, "", "")
	restErr := uc.userService.GetAccount(userDomain, cookie.SessionId)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	photo, restErr := user_service_util.GetUserImage(userDomain.GetUser())
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(http.StatusOK, user_response.GetAccount{
		Name: userDomain.GetName(),
		User: userDomain.GetUser(),
		Category: userDomain.GetCategory(),
		Earnings: userDomain.GetEarnings(),
		Photo: photo,
	})
}