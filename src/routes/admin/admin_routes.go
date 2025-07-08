package admin_router

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
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

	r.Use(func(c *gin.Context) {
		_, err := admin_profile_cookie.GetAdminProfileValues(c)
		if err != nil {
			logger.Error("Error trying get cookie", err, zap.String("journey", "admin route"))
			restErr := rest_err.NewUnauthorizedError("cookie inválido")
			c.JSON(restErr.Code, restErr)
			c.Abort()
			return
		}
		c.Next()
	})

	r.POST("/manager-quikworkout/createEdition", adminController.CreateEdition)
	r.GET("/manager-quikworkout/getParticipants", adminController.GetParticipants)
	r.POST("/manager-quikworkout/checkVideo", adminController.CheckVideo)
	r.POST("/manager-quikworkout/desqualifyVideo", adminController.DesqualifyVideo)
	r.POST("/manager-quikworkout/makePlacing", adminController.MakePlacing)
	r.POST("/manager-quikworkout/putTime", adminController.PutTime)
	r.POST("/manager-quikworkout/grantTicket", adminController.GrantTicket)
}

func initAdminRoutes(database *sql.DB) admin_controller.AdminController {
	adminRepository := admin_repository.NewAdminRepository(database)
	adminService := admin_service.NewAdminService(adminRepository)
	adminController := admin_controller.NewAdminController(adminService)
	return adminController
}
