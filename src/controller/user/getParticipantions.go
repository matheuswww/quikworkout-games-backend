package user_controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	user_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/request"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	model_util "github.com/matheuswww/quikworkout-games-backend/src/model/util"
	"go.uber.org/zap"
)

func (uc *userController) GetParticipations(c *gin.Context) {
	cookie, err := user_games_cookie.GetUserGamesCookieValues(c)
	if err != nil {
		logger.Error("Erro trying get cookie", err, zap.String("journey", "GetParticipants Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}
	restErr := model_util.CheckUserGames(cookie.SessionId, cookie.Id)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	var getParticipationsRequest user_request.GetParticipations
	if err := c.ShouldBind(&getParticipationsRequest); err != nil {
		logger.Error("Error trying convert fields", errors.New("error trying convert fields"), zap.String("journey", "GetParticipants Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		if restErr != nil {
			c.JSON(restErr.Code, restErr)
			return
		}
	}

	userDomain := user_domain.NewUserDomain(cookie.Id, "", "", "", 0, "", "")
	participations, restErr := uc.userService.GetParticipations(userDomain, &getParticipationsRequest)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(http.StatusOK, participations)
}