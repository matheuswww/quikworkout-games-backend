package participant_controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	participant_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/request"
	participant_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/response"
	"go.uber.org/zap"
)

func (pc *participantController) GetParticipants(c *gin.Context) {
	logger.Info("Init GetParticiapant", zap.String("journey", "GetParticipant Repository"))
	var getParticipartRequest participant_request.GetParticipant
	if err := c.ShouldBindQuery(&getParticipartRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("error trying convert fields"), zap.String("journey", "GetParticipant Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	participants, restErr := pc.participantService.GetParticipants(getParticipartRequest.EditionId, getParticipartRequest.CursorCreatedAt, getParticipartRequest.CursorUserTime, getParticipartRequest.WorstTime)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, participant_response.GetParticipant{
		Particiapants: participants,
	})
}