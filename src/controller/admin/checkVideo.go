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

func (ac *adminController) CheckVideo(c *gin.Context) {
	logger.Info("Init CheckVideo Controller", zap.String("journey", "CheckVideo Controller"))
	var checkVideoRequest admin_request.CheckVideo
	if err := c.ShouldBindJSON(&checkVideoRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("error trying convert fields"), zap.String("journey", "ChckVideo Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(checkVideoRequest)
	restErr := ac.adminService.CheckVideo(checkVideoRequest.VideoID, checkVideoRequest.EditionId, checkVideoRequest.Category, checkVideoRequest.Sex)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("Video checekd! User_id: "+checkVideoRequest.VideoID, zap.String("journey", "CheckVideo Controller"))
	c.Status(http.StatusOK)
}