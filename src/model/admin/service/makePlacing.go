package admin_service

import "github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"

func (as *adminService) MakePlacing(editionId, category string) *rest_err.RestErr {
	return as.adminRepository.MakePlacing(editionId, category)
}