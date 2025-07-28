package participant_controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	custom_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/customValidator"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	get_custom_validator "github.com/matheuswww/quikworkout-games-backend/src/controller/model"
	participant_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/request"
	participant_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/response"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	participant_domain "github.com/matheuswww/quikworkout-games-backend/src/model/participant"
	model_util "github.com/matheuswww/quikworkout-games-backend/src/model/util"
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
	translator, customErr := get_custom_validator.CustomValidator(createParticipantRequest)
	if customErr != nil {
		restErr := custom_validator.HandleCustomValidatorErrors(translator, customErr)
		logger.Error("Error trying convert fields", errors.New("invalid fields"), zap.String("journey", "CreateEdition Controller"))
		c.JSON(restErr.Code, restErr)
		return
	}

	participant := participant_domain.NewParticipantDomain("", cookie.Id, "", "", "", createParticipantRequest.UserTime, "", createParticipantRequest.Sex, false, false)

	form, restErr := pc.participantService.CreateParticipant(participant, createParticipantRequest.Title, createParticipantRequest.Size)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("Participant created! user_id: "+cookie.Id, zap.String("journey", "CreateParticipant Controller"))
	c.JSON(http.StatusCreated, participant_response.CreateParticipant{
		Form: form,
	})
}