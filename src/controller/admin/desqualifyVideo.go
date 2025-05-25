package admin_controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	admin_profile_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/admin_profile"
	"go.uber.org/zap"
)

func (ac *adminController) DesqualifyVideo(c *gin.Context) {
	logger.Info("Init DesqualifyVideo", zap.String("journey", "DesqualifyVideo Controller"))
	_, err := admin_profile_cookie.GetAdminProfileValues(c)
	if err != nil {
		logger.Error("Erro trying get cookie", err, zap.String("journey", "DesqualifyVideo Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}

	var desqualifyVideoRequest admin_request.DesqualifyVideo
	if err := c.ShouldBindJSON(&desqualifyVideoRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("error trying convert fields"), zap.String("journey", "DesqualifyVideo Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	restErr := ac.adminService.DesqualifyVideo(desqualifyVideoRequest.VideoID, desqualifyVideoRequest.Desqualified)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("Video desqualified! video_id: "+desqualifyVideoRequest.VideoID, zap.String("journey", "DesqualifyVideo Controller"))
	c.Status(http.StatusOK)
}