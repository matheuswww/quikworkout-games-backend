package admin_service

import "github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"

func (as *adminService) DesqualifyVideo(videoID, desqualifed string) *rest_err.RestErr {
	return as.adminRepository.DesqualifyVideo(videoID, desqualifed)
}