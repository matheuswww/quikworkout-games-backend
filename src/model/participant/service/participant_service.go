package participant_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	participant_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/request"
	participant_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/response"
	participant_domain "github.com/matheuswww/quikworkout-games-backend/src/model/participant"
	participant_repository "github.com/matheuswww/quikworkout-games-backend/src/model/participant/repository"
)

type participantService struct {
	participantRepository participant_repository.ParticipantRepository
}

func NewParticipantService(participantRepository participant_repository.ParticipantRepository) ParticipantService {
	return &participantService{
		participantRepository,
	}
}

type ParticipantService interface {
	CreateParticipant(participantDomain participant_domain.ParticipantDomainInterface, title string, size int64) (string, *rest_err.RestErr)
	GetParticipants(getParticipantRequest *participant_request.GetParticipant) (*participant_response.GetParticipant, *rest_err.RestErr)
	HasTicket(cookieId string) *rest_err.RestErr
	VideoSent(videoId, userId string) *rest_err.RestErr
}
