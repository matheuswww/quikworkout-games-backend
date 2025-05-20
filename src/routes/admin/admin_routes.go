package admin_router

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	admin_controller "github.com/matheuswww/quikworkout-games-backend/src/controller/admin"
	admin_profile_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/admin_profile"
	admin_repository "github.com/matheuswww/quikworkout-games-backend/src/model/admin/repository"
	admin_service "github.com/matheuswww/quikworkout-games-backend/src/model/admin/service"
	"go.uber.org/zap"
)

func InitAdminRoutes(r *gin.RouterGroup, database *sql.DB) {
	adminController := initAdminRoutes(database)
	cookieStore, err := admin_profile_cookie.Store()
	if err != nil {
		logger.Error("Error loading cookie store", err, zap.String("journey", "InitAdminRoutes"))
		log.Fatal("Error cookie store")
	}
	sessionNames := []string{admin_profile_cookie.SessionAdminProfile}
	r.Use(sessions.SessionsMany(sessionNames, cookieStore))

	r.POST("/manager-quikworkout/createEdition", adminController.CreateEdition)
}

func initAdminRoutes(database *sql.DB) admin_controller.AdminController {
	adminRepository := admin_repository.NewAdminRepository(database)
	adminService := admin_service.NewAdminService(adminRepository)
	adminController := admin_controller.NewAdminController(adminService)
	return adminController
}
