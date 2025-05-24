package participant_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/vimeo"
	participant_domain "github.com/matheuswww/quikworkout-games-backend/src/model/participant"
)

func (ps *participantService) CreateParticipant(participantDomain participant_domain.ParticipantDomainInterface, title, instagram string, size int64) (string, *rest_err.RestErr) {
	restErr := ps.participantRepository.IsValidRegistrationForEdition(participantDomain)
	if restErr != nil {
		return "", restErr
	}
	form, id, err := vimeo.UploadVideo(title, size)
	if err != nil {
		restErr := rest_err.NewInternalServerError("server error")
		return "", restErr
	}
	participantDomain.SetVideoID(id)
	restErr = ps.participantRepository.CreateParticipant(participantDomain, instagram)
	if restErr != nil {
		return "", restErr
	}
	return form, nil
}