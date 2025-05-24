package participant_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
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
	CreateParticipant(participantDomain participant_domain.ParticipantDomainInterface, title, instagram string, size int64) (string, *rest_err.RestErr)
	GetParticipant(editionID, cursor_createdAt, cursor_userTime string, worstTime bool) ([]participant_response.Participant, *rest_err.RestErr)
}
