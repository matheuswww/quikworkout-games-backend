package admin_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	admin_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/response"
)

func (as *adminService) GetParticipants(editionID, cursor_createdAt, cursor_userTime string) ([]admin_response.Participant, *rest_err.RestErr) {
	return as.adminRepository.GetParticipants(editionID, cursor_createdAt, cursor_userTime)
}