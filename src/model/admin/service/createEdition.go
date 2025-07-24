package admin_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
)

func (as *adminService) CreateEdition(createEditionRequest *admin_request.CreateEdition, tops []admin_request.Top, challenges []admin_request.Challenge, savePdf func(id string) *rest_err.RestErr) (*rest_err.RestErr) {
	return as.adminRepository.CreateEdition(createEditionRequest, tops, challenges, savePdf)
}