package admin_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	custom_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/customValidator"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	admin_controller_util "github.com/matheuswww/quikworkout-games-backend/src/controller/admin/util"
	get_custom_validator "github.com/matheuswww/quikworkout-games-backend/src/controller/model"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	"go.uber.org/zap"
)

func (ac *adminController) CreateEdition(c *gin.Context) {
	logger.Info("Init CreateEdition", zap.String("journey", "CreateEdition Controller"))
	var createEdition admin_request.CreateEdition
	if err := c.ShouldBind(&createEdition); err != nil {
		logger.Error("Error trying convert fileds", errors.New("invalid fields"), zap.String("journey", "CreateEdition controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	translator, customErr := get_custom_validator.CustomValidator(createEdition)
	if customErr != nil {
		restErr := custom_validator.HandleCustomValidatorErrors(translator, customErr)
		logger.Error("Error trying convert fields", errors.New("invalid fields"), zap.String("journey", "CreateEdition Controller"))
		c.JSON(restErr.Code, restErr)
		return
	}

	var challengesRequest []admin_request.Challenge
	if err := json.Unmarshal([]byte(createEdition.Challenge), &challengesRequest); err != nil {
		logger.Error("Error unmarshaling challenge JSON", err, zap.String("journey", "CreateEdition controller"))
		restErr := rest_err.NewBadRequestError("invalid challenges")
		c.JSON(restErr.Code, restErr)
		return
	}

	var topsRequest []admin_request.Top
	if err := json.Unmarshal([]byte(createEdition.Tops), &topsRequest); err != nil {
		logger.Error("Error unmarshaling tops JSON", err, zap.String("journey", "CreateEdition controller"))
		restErr := rest_err.NewBadRequestError("invalid tops")
		c.JSON(restErr.Code, restErr)
		return
	}
	restErr := admin_controller_util.ValidateChallengeAndTop(challengesRequest, topsRequest)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	if !strings.HasSuffix(strings.ToLower(createEdition.Rules.Filename), ".pdf") {
    restErr := rest_err.NewBadRequestError("file must be pdf")
		c.JSON(restErr.Code, restErr)
		return
	}

	tops := make(map[string]map[int]bool)
	lastCategory := topsRequest[0].Category
	i := 0
	for _,top := range topsRequest {
		if lastCategory != top.Category {
			i = 0
		}
		i++
		if ok := tops[top.Category][top.Top]; ok {
			restErr := rest_err.NewBadRequestError("não é permitido ter dois tops iguais")
			c.JSON(restErr.Code, restErr)
			return
		}
		if i != top.Top {
			restErr := rest_err.NewBadRequestError("os tops devem ser ordenados em 1,2,3...")
			c.JSON(restErr.Code, restErr)
			return
		}
		if tops[top.Category] == nil {
			tops[top.Category] = make(map[int]bool)
		}
	
		tops[top.Category][top.Top] = true
		lastCategory = top.Category
	}
	
	restErr = ac.adminService.CreateEdition(&createEdition,topsRequest, challengesRequest, func(id string) *rest_err.RestErr {
		absPath, err := filepath.Abs("pdf")
		if err != nil {
			logger.Error("Error trying saveImg", err, zap.String("journey", "SavePdf"))
			return rest_err.NewInternalServerError("server error")
		}
		err = c.SaveUploadedFile(createEdition.Rules, fmt.Sprintf("%s/%s.pdf", absPath, id))
		if err != nil {
			logger.Error("Error trying saveUploadedFile", err, zap.String("journey", "SavePdf"))
			return rest_err.NewInternalServerError("server error")
		}
		return nil
	})
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	logger.Info("Edition created!", zap.String("journey", "CreateEdition Controller"))
	c.Status(http.StatusCreated)
}