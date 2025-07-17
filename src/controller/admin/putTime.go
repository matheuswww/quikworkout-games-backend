package admin_controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	"go.uber.org/zap"
)

func (ac *adminController) PutTime(c *gin.Context) {
	logger.Info("Init PutTime Controller", zap.String("journey", "PutTime Controller"))
	var putTimeRequest admin_request.PutTimeRequest
	if err := c.ShouldBindJSON(&putTimeRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("invalid fields"), zap.String("journey", "PutTime controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	restErr := ac.adminService.PutTime(putTimeRequest.VideoId, putTimeRequest.EditionId, putTimeRequest.Category, putTimeRequest.Sex, putTimeRequest.Time)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info(fmt.Sprintf("Time puted with success, video_id: %s", putTimeRequest.VideoId), zap.String("journey", "PutTime Controller"))
	c.Status(http.StatusOK)
}