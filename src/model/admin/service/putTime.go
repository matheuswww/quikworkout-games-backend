package admin_service

import "github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"

func (as *adminService) PutTime(videoId, editionId, category, sex, userTime string) *rest_err.RestErr {
	return as.adminRepository.PutTime(videoId, editionId, category, sex, userTime)
}