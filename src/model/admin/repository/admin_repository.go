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
	GetParticipants(editionID, cursor_createdAt, cursor_userTime string) ([]admin_response.Participant, *rest_err.RestErr)
}
