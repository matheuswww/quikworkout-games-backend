package user_controller

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
	user_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/request"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	user_proflie_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_profile"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	"go.uber.org/zap"
)

func (uc *userController) CreateAccount(c *gin.Context) {
	logger.Info("Init CreateAccount")
	var createAccountRequest user_request.CreateAccount
	if err := c.ShouldBindJSON(&createAccountRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("invalid fields"), zap.String("journey", "CreateAccount Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	translator, customErr := get_custom_validator.CustomValidator(createAccountRequest)
	if customErr != nil {
		restErr := custom_validator.HandleCustomValidatorErrors(translator, customErr)
		logger.Error("Error trying convert fields", errors.New("invalid fields"), zap.String("journey", "CreateAccount Controller"))
		c.JSON(restErr.Code, restErr)
		return
	}
	cookie, err := user_proflie_cookie.GetUserProfileCookieValues(c)
	if err != nil {
		logger.Error("Erro trying get cookie", err, zap.String("journey", "PayOrder Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
	}
	userDomain := user_domain.NewUserDomain(cookie.Id, "",  createAccountRequest.User, createAccountRequest.DOB, createAccountRequest.Category, 0, createAccountRequest.CPF, "")
	restErr := uc.userService.CreateAccount(userDomain, cookie.SessionId, createAccountRequest.Token)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	cookieErr := user_games_cookie.SendUserGamesCookie(c, userDomain.GetId(), userDomain.GetSessionId(), true)
	if cookieErr != nil {
		logger.Error("Error trying create session", cookieErr, zap.String("journey", "CreateAccount Controller"))
		restErr := rest_err.NewInternalServerError("usuário criado porém não foi possível criar uma sessão")
		c.JSON(restErr.Code, restErr)
		return
	}
	c.Header("Access-Control-Expose-Headers", "Set-Cookie")
	logger.Info(fmt.Sprintf("User registred with success!,id: %s", userDomain.GetId()), zap.String("journey", "CreateAccount Controller"))
	c.Status(http.StatusCreated)
}