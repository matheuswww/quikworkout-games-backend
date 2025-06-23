package participant_repository

import (
	"database/sql"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	participant_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/request"
	participant_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/response"
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
	CreateParticipant(participantDomain participant_domain.ParticipantDomainInterface) *rest_err.RestErr 
	IsValidRegistrationForEdition(participantDomain participant_domain.ParticipantDomainInterface) *rest_err.RestErr
	GetParticipants(getParticipantRequest *participant_request.GetParticipant) (*participant_response.GetParticipant, *rest_err.RestErr)
	HasTicket(cookieId string) ([]PaymentInfos, *rest_err.RestErr)
	VideoSent(videoId, userId string) *rest_err.RestErr
}
