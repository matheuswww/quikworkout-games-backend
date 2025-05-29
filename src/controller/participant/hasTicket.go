package participant_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	"go.uber.org/zap"
)

func (ec *participantController) HasTicket(c *gin.Context) {
	logger.Info("GetTicket Controller", zap.String("journey", "GetTicket Controller"))

	cookie, err := user_games_cookie.GetUserGamesCookieValues(c)
	if err != nil {
		logger.Error("Error trying get cookie", err, zap.String("journey", "CreateParticiapant Controller Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}

	restErr := ec.participantService.HasTicket(cookie.Id)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusOK)
}