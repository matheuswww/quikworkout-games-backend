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

func (ac *adminController) PutNoReps(c *gin.Context) {
	logger.Info("Init PutNoReps Controller", zap.String("journey", "PutNoReps Controller"))
	var putNoRepsRequest admin_request.PutNoreps
	if err := c.ShouldBindJSON(&putNoRepsRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("invalid fields"), zap.String("journey", "PutNoReps controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	restErr := ac.adminService.PutNoreps(&putNoRepsRequest)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info(fmt.Sprintf("No reps puted with success, video_id: %s", putNoRepsRequest.VideoId), zap.String("journey", "PutNoReps Controller"))
	c.Status(http.StatusOK)
}