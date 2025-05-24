package participant_controller

import (
	"github.com/gin-gonic/gin"
	participant_service "github.com/matheuswww/quikworkout-games-backend/src/model/participant/service"
)

type participantController struct {
	participantService participant_service.ParticipantService
}

func NewParticipantController(participantService participant_service.ParticipantService) ParticipantController {
	return &participantController{
		participantService,
	}
}

type ParticipantController interface {
	CreateParticipant(c *gin.Context)
	GetParticipant(c *gin.Context)
}
