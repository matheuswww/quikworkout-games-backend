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

func (ac *adminController) GrantTicket(c *gin.Context) {
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