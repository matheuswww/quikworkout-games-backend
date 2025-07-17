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
	"go.uber.org/zap"
)

func (pc *participantController) GetParticipants(c *gin.Context) {
	logger.Info("Init GetParticiapant", zap.String("journey", "GetParticipant Repository"))
	var getParticipantRequest participant_request.GetParticipant
	if err := c.ShouldBindQuery(&getParticipantRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("error trying convert fields"), zap.String("journey", "GetParticipant Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	translator, customErr := get_custom_validator.CustomValidator(getParticipantRequest)
	if customErr != nil {
		restErr := custom_validator.HandleCustomValidatorErrors(translator, customErr)
		logger.Error("Error trying convert fields", errors.New("invalid fields"), zap.String("journey", "GetParticipant Controller"))
		c.JSON(restErr.Code, restErr)
		return
	}
	if getParticipantRequest.Category == "" && getParticipantRequest.VideoId == "" {
		logger.Error("Error trying get participants", errors.New("category or video_id must be provided"), zap.String("journey", "GetParticipant Repository"))
		restErr := rest_err.NewBadRequestError("category or video_id must be provided")
		c.JSON(restErr.Code, restErr);
		return
	}
	if getParticipantRequest.Category != "" || getParticipantRequest.Sex != "" {
		if getParticipantRequest.Category == "" || getParticipantRequest.Sex == "" {
			logger.Error("Error trying get participants", errors.New("category and sex must be provided"), zap.String("journey", "GetParticipant Repository"))
		restErr := rest_err.NewBadRequestError("category and sex must be provided")
		c.JSON(restErr.Code, restErr);
		return
		}
	}


	participants, restErr := pc.participantService.GetParticipants(&getParticipantRequest)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, participants)
}