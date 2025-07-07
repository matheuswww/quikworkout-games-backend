package admin_controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	custom_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/customValidator"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	get_custom_validator "github.com/matheuswww/quikworkout-games-backend/src/controller/model"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	admin_profile_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/admin_profile"
	"go.uber.org/zap"
)

func (ac *adminController) MakePlacing(c *gin.Context) {
	logger.Info("Init MakePlacing Controller")
	_, err := admin_profile_cookie.GetAdminProfileValues(c)
	if err != nil {
		logger.Error("Erro trying get cookie", err, zap.String("journey", "MakePlacing Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}
	var makePlacingRequest admin_request.MakePlacing
	if err := c.ShouldBindJSON(&makePlacingRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("invalid fields"), zap.String("journey", "MakePlacing controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	translator, customErr := get_custom_validator.CustomValidator(makePlacingRequest)
	if customErr != nil {
		restErr := custom_validator.HandleCustomValidatorErrors(translator, customErr)
		logger.Error("Error trying convert fields", errors.New("invalid fields"), zap.String("journey", "MakePlacing Controller"))
		c.JSON(restErr.Code, restErr)
		return
	}


	restErr := ac.adminService.MakePlacing(makePlacingRequest.EditionId, makePlacingRequest.Category)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info(fmt.Sprintf("Placing made with success, edition_id: %s", makePlacingRequest.EditionId), zap.String("journey", "MakePlacing Controller"))
	c.Status(http.StatusOK)
}