package admin_repository

import (
	"database/sql"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	admin_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/response"
)

func NewAdminRepository(mysql *sql.DB) AdminRepository {
	return &adminRepository{
		mysql,
	}
}

type adminRepository struct {
	mysql *sql.DB
}

type AdminRepository interface {
	CreateEdition(createEditionRequest *admin_request.CreateEdition) *rest_err.RestErr
	GetParticipants(getParticipantsRequest *admin_request.GetParticipants) ([]admin_response.Participant, *sql.DB, *rest_err.RestErr)
	CheckVideo(videoID string) *rest_err.RestErr
	DesqualifyVideo(videoID, desqualifed string) *rest_err.RestErr
	MakePlacing(editionId string) *rest_err.RestErr
}
