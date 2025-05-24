package participant_repository

import (
	"database/sql"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	participant_domain "github.com/matheuswww/quikworkout-games-backend/src/model/participant"
)

type participantRepository struct {
	mysql *sql.DB
}

func NewParticipantRepository(mysql *sql.DB) ParticipantRepository {
	return &participantRepository{
		mysql,
	}
}

type ParticipantRepository interface {
	CreateParticipant(participantDomain participant_domain.ParticipantDomainInterface, instagram string) *rest_err.RestErr 
	IsValidRegistrationForEdition(participantDomain participant_domain.ParticipantDomainInterface) *rest_err.RestErr
}
