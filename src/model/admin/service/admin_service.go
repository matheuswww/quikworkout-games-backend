package admin_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	admin_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/response"
	admin_repository "github.com/matheuswww/quikworkout-games-backend/src/model/admin/repository"
)

func NewAdminService(adminRepository admin_repository.AdminRepository) AdminService {
	return &adminService{
		adminRepository,
	}
}

type adminService struct {
	adminRepository admin_repository.AdminRepository
}

type AdminService interface {
	CreateEdition(createEditionRequest *admin_request.CreateEdition) *rest_err.RestErr
	GetParticipants(editionID, cursor_createdAt, cursor_userTime string) ([]admin_response.Participant, *rest_err.RestErr)
	CheckVideo(videoID string) *rest_err.RestErr
	DesqualifyVideo(videoID, desqualifed string) *rest_err.RestErr
	MakePlacing(editionId string) *rest_err.RestErr
}
