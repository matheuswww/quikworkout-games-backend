package user_controller

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	custom_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/customValidator"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	get_custom_validator "github.com/matheuswww/quikworkout-games-backend/src/controller/model"
	user_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/request"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	model_util "github.com/matheuswww/quikworkout-games-backend/src/model/util"
	"go.uber.org/zap"
)

func (uc *userController) Update(c *gin.Context) {
	logger.Info("Init Update")
	cookie, err := user_games_cookie.GetUserGamesCookieValues(c)
	if err != nil {
		logger.Error("Erro trying get cookie", err, zap.String("journey", "Update Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}
	restErr := model_util.CheckUserGames(cookie.SessionId, cookie.Id)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	var updateRequest user_request.Update
	if err := c.ShouldBind(&updateRequest); err != nil {
		logger.Error("Error trying convert fileds", errors.New("invalid fields"), zap.String("journey", "Update Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	translator, customErr := get_custom_validator.CustomValidator(updateRequest)
	if customErr != nil {
		restErr := custom_validator.HandleCustomValidatorErrors(translator, customErr)
		logger.Error("Error trying convert fields", errors.New("invalid fields"), zap.String("journey", "Update Controller"))
		c.JSON(restErr.Code, restErr)
		return
	}
	if updateRequest.User == "" && updateRequest.Category == "" && updateRequest.Name == "" && updateRequest.Image == nil {
		restErr := rest_err.NewBadRequestError("invalid params")
		c.JSON(restErr.Code, restErr)
		return
	}
	userDomain := user_domain.NewUserDomain(cookie.Id, "", "", "", 0, "")
	if updateRequest.Image != nil {
		err = saveNewImg(c, &updateRequest, userDomain.GetId())
		if err != nil {
			return
		}
	}
	if updateRequest.User != "" || updateRequest.Name != "" || updateRequest.Category != "" {
		restErr = uc.userService.Update(userDomain, &updateRequest)
		if restErr != nil {
			c.JSON(restErr.Code, restErr)
			return
		}
	}
	c.Status(http.StatusOK)
}

func saveNewImg(c *gin.Context, updateRequest *user_request.Update, id string) error {
	const maxSize = 1 * 1024 * 1024
		if updateRequest.Image.Size > maxSize {
			restErr := rest_err.NewBadRequestError("image size must be less than or equal to 1MB")
			c.JSON(restErr.Code, restErr)
			return errors.New("invalid size")
		}
		ext := strings.ToLower(filepath.Ext(updateRequest.Image.Filename))
		allowedExts := map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		}
		if !allowedExts[ext] {
			restErr := rest_err.NewBadRequestError("image extension must be jpg, jpeg, or png")
			c.JSON(restErr.Code, restErr)
			return errors.New("invalid extension")
		}

		absPath, err := filepath.Abs("images")
		if err != nil {
			logger.Error("Error trying update image", err, zap.String("journey", "Update Controller"))
			restErr := rest_err.NewInternalServerError("server error")
			c.JSON(restErr.Code, restErr)
			return err
		}
		
		files, err := os.ReadDir(absPath)
		if err != nil {
			logger.Error("Error reading images directory", err, zap.String("journey", "Update Controller"))
			restErr := rest_err.NewInternalServerError("server error")
			c.JSON(restErr.Code, restErr)
			return err
		}
	
		for _, file := range files {
			if strings.HasPrefix(file.Name(), id) {
				err = os.Remove(filepath.Join(absPath, file.Name()))
				if err != nil {
					logger.Error("Error deleting old image", err, zap.String("journey", "Update Controller"))
					restErr := rest_err.NewInternalServerError("server error")
					c.JSON(restErr.Code, restErr)
					return err
				}
				break
			}
		}

		ext = filepath.Ext(updateRequest.Image.Filename)
		err = c.SaveUploadedFile(updateRequest.Image, fmt.Sprintf("%s/%s%s", absPath, id, ext))
		if err != nil {
			logger.Error("Error trying get update image", err, zap.String("journey", "Update Controller"))
			restErr := rest_err.NewInternalServerError("server error")
			c.JSON(restErr.Code, restErr)
			return err
		}
		return nil
}