package user_controller

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/sightengine"
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
	cookie, err := user_proflie_cookie.GetUserProfileCookieValues(c)
	if err != nil {
		logger.Error("Erro trying get cookie", err, zap.String("journey", "CreateAccount Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}
	var createAccountRequest user_request.CreateAccount
	if err := c.ShouldBind(&createAccountRequest); err != nil {
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
	const maxSize = 3 * 1024 * 1024
	if createAccountRequest.Image.Size > maxSize {
		restErr := rest_err.NewBadRequestError("image size must be less than or equal to 3MB")
		c.JSON(restErr.Code, restErr)
		return
	}
	ext := strings.ToLower(filepath.Ext(createAccountRequest.Image.Filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	if !allowedExts[ext] {
		restErr := rest_err.NewBadRequestError("image extension must be jpg, jpeg, or png")
		c.JSON(restErr.Code, restErr)
		return
	}
	
	userDomain := user_domain.NewUserDomain(cookie.Id, "", createAccountRequest.User, createAccountRequest.Category, 0, "")
	restErr := uc.userService.CreateAccount(userDomain, func() *rest_err.RestErr {
		absPath, err := filepath.Abs("images")
		if err != nil {
			logger.Error("Error trying saveImg", err, zap.String("journey", "CreateAccount Controller"))
			return rest_err.NewInternalServerError("server error")
		}
		restErr := sightengine.CheckImage(createAccountRequest.Image, fmt.Sprintf("%s%s", userDomain.GetId(), ext))
		if restErr != nil {
			return restErr
		}
		ext := filepath.Ext(createAccountRequest.Image.Filename)
		err = c.SaveUploadedFile(createAccountRequest.Image, fmt.Sprintf("%s/%s%s", absPath, userDomain.GetId(), ext))
		if err != nil {
			logger.Error("Error trying saveImg", err, zap.String("journey", "CreateAccount Controller"))
			return rest_err.NewInternalServerError("server error")
		}
		return nil
	} , cookie.SessionId, createAccountRequest.Token)
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