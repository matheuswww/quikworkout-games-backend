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
	CreateEdition(createEditionRequest *admin_request.CreateEdition, tops []admin_request.Top, challenges []admin_request.Challenge, savePdf func(id string) *rest_err.RestErr) (*rest_err.RestErr)
	GetParticipants(getParticipantsRequest *admin_request.GetParticipants) (*admin_response.GetParticipants, *rest_err.RestErr)
	CheckVideo(videoID, category, sex, editionId string) *rest_err.RestErr
	DesqualifyVideo(videoID, editionId, category, sex, desqualifed string) *rest_err.RestErr
	MakePlacing(editionId, category, sex string) *rest_err.RestErr
	PutTime(videoId, editionId, category, sex, finalTime string) *rest_err.RestErr
	GrantTicket(user string) *rest_err.RestErr
	PutNoreps(putNorepsRequest *admin_request.PutNoreps) *rest_err.RestErr 
}