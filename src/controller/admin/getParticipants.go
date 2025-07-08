package admin_controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	custom_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/customValidator"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	get_custom_validator "github.com/matheuswww/quikworkout-games-backend/src/controller/model"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	"go.uber.org/zap"
)

func (ac *adminController) GetParticipants(c *gin.Context) {
	logger.Info("Init GetParticipants Controller", zap.String("journey", "GetParticipants Controller"))
	var getParticipantsRequest admin_request.GetParticipants
	if err := c.ShouldBindQuery(&getParticipantsRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("invalid fields"), zap.String("journey", "GetParticipants controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	translator, customErr := get_custom_validator.CustomValidator(getParticipantsRequest)
	if customErr != nil {
		restErr := custom_validator.HandleCustomValidatorErrors(translator, customErr)
		logger.Error("Error trying convert fields", errors.New("invalid fields"), zap.String("journey", "GetParticipants Controller"))
		c.JSON(restErr.Code, restErr)
		return
	}

	participants, restErr := ac.adminService.GetParticipants(&getParticipantsRequest)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, participants)
}