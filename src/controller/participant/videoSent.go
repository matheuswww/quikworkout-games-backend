package participant_controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	participant_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/request"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	model_util "github.com/matheuswww/quikworkout-games-backend/src/model/util"
	"go.uber.org/zap"
)

func (pc *participantController) VideoSent(c *gin.Context) {
	logger.Info("Init VideoSent Controller", zap.String("journey", "VideoSent Controller"))
	cookie, err := user_games_cookie.GetUserGamesCookieValues(c)
	if err != nil {
		logger.Error("Error trying get cookie", err, zap.String("journey", "VideoSent Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}
	restErr := model_util.CheckUserGames(cookie.SessionId, cookie.Id)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	var videoSentRequest participant_request.VideoSent
	if err := c.ShouldBindJSON(&videoSentRequest); err != nil {
		logger.Error("Error trying convert fields", err, zap.String("journey", "VideoSent Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	restErr = pc.participantService.VideoSent(videoSentRequest.VideoId, cookie.Id)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info(fmt.Sprintf("Video sent updated with successs, video_id: %s", videoSentRequest.VideoId), zap.String("journey", "VideoSent Controller"))
	c.Status(http.StatusOK)
}