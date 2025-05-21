package admin_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
)

func (as *adminService) CreateEdition(createEditionRequest *admin_request.CreateEdition) *rest_err.RestErr {
	tops := make(map[int]bool)
	for _,top := range createEditionRequest.Tops {
		if ok := tops[top.Top]; ok {
			return rest_err.NewBadRequestError("não é permitido ter dois tops iguais")
		}
		tops[top.Top] = true
	}
	return as.adminRepository.CreateEdition(createEditionRequest)
}