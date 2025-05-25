package participant_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	participant_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/response"
)

func (ps *participantService) GetParticipants(editionID, cursor_createdAt, cursor_userTime string, worstTime bool) ([]participant_response.Participant, *rest_err.RestErr) {
	return ps.participantRepository.GetParticipants(editionID, cursor_createdAt, cursor_userTime, worstTime)
}