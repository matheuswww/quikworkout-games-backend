package user_controller

import (
	"github.com/gin-gonic/gin"
	user_service "github.com/matheuswww/quikworkout-games-backend/src/model/user/service"
)

func NewUserController(userService user_service.UserService) UserController {
	return &userController{
		userService,
	}
}

type userController struct {
	userService user_service.UserService
}

type UserController interface {
	CreateAccount(c *gin.Context)
	EnterAccount(c *gin.Context)
	GetAccount(c *gin.Context)
	GetParticipations(c *gin.Context)
}
