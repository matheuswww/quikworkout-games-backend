package admin_controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	"go.uber.org/zap"
)

func (ac *adminController) DesqualifyVideo(c *gin.Context) {
	logger.Info("Init DesqualifyVideo", zap.String("journey", "DesqualifyVideo Controller"))
	var desqualifyVideoRequest admin_request.DesqualifyVideo
	if err := c.ShouldBindJSON(&desqualifyVideoRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("error trying convert fields"), zap.String("journey", "DesqualifyVideo Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	restErr := ac.adminService.DesqualifyVideo(desqualifyVideoRequest.VideoID, desqualifyVideoRequest.EditionId, desqualifyVideoRequest.Desqualified)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("Video desqualified! video_id: "+desqualifyVideoRequest.VideoID, zap.String("journey", "DesqualifyVideo Controller"))
	c.Status(http.StatusOK)
}