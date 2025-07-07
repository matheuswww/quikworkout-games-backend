package admin_controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	admin_profile_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/admin_profile"
	"go.uber.org/zap"
)

func (ac *adminController) GrantTicket(c *gin.Context) {
	_, err := admin_profile_cookie.GetAdminProfileValues(c)
	if err != nil {
		logger.Error("Erro trying get cookie", err, zap.String("journey", "GrantTicket Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}
	var grantTicket admin_request.GrantTicket
	if err := c.ShouldBindJSON(&grantTicket); err != nil {
		logger.Error("Error trying convert fileds", errors.New("invalid fields"), zap.String("journey", "GrantTicket controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	restErr := ac.adminService.GrantTicket(grantTicket.User)
	if restErr != nil {
		logger.Error("Error trying GrantTicket", restErr, zap.String("journey", "GrantTicket Controller"))
		c.JSON(restErr.Code, restErr)
		return
	}
	
	logger.Info(fmt.Sprintf("Ticket granted with sucess, user: %s", grantTicket.User), zap.String("journey", "GrantTicket Controller"))
	c.Status(http.StatusOK)
}