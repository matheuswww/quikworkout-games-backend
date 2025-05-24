package participant_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	participant_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/request"
	participant_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/response"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	model_util "github.com/matheuswww/quikworkout-games-backend/src/model/util"
	participant_domain "github.com/matheuswww/quikworkout-games-backend/src/model/participant"
	"go.uber.org/zap"
)

func (pc *participantController) CreateParticipant(c *gin.Context) {
	logger.Info("Init CreateParticipant", zap.String("journey", "CreateParticipant Controller"))
	
	cookie, err := user_games_cookie.GetUserGamesCookieValues(c)
	if err != nil {
		logger.Error("Error trying get cookie", err, zap.String("journey", "CreateParticiapant Controller Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}
	restErr := model_util.CheckUserGames(cookie.SessionId, cookie.Id)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	var createParticipantRequest participant_request.CreateParticipant
	if err := c.ShouldBindJSON(&createParticipantRequest); err != nil {
		logger.Error("Error trying convert fields", err, zap.String("journey", "CreateParticipant Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	participant := participant_domain.NewParticipantDomain("", cookie.Id, "", nil, "", false)

	form, restErr := pc.participantService.CreateParticipant(participant, createParticipantRequest.Title, createParticipantRequest.Instagram, createParticipantRequest.Size)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, participant_response.CreateParticipant{
		Form: form,
	})
}