package admin_controller

import (
	"github.com/gin-gonic/gin"
	admin_service "github.com/matheuswww/quikworkout-games-backend/src/model/admin/service"
)


func NewAdminController(adminService admin_service.AdminService) AdminController {
	return &adminController{
		adminService,
	}
}

type adminController struct {
	adminService admin_service.AdminService
}

type AdminController interface {
	CreateEdition(c *gin.Context)
}