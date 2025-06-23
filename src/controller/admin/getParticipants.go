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

func (ac *adminController) GetParticipants(c *gin.Context) {
	logger.Info("Init GetParticipants Controller", zap.String("journey", "GetParticipants Controller"))
	_, err := admin_profile_cookie.GetAdminProfileValues(c)
	if err != nil {
		logger.Error("Error trying get cookie", err, zap.String("journey", "GetParticipants Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}
	var getParticipantsRequest admin_request.GetParticipants
	if err := c.ShouldBindQuery(&getParticipantsRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("invalid fields"), zap.String("journey", "GetParticipants controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
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