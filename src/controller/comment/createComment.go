package comment_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	comment_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/comment/request"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	comment_domain "github.com/matheuswww/quikworkout-games-backend/src/model/comment"
	"go.uber.org/zap"
)

func (cc *commentController) CreateComment(c *gin.Context) {
	logger.Info("Init CreateComment controller", zap.String("journey", "CreateComment Controller"))

	cookie, err := user_games_cookie.GetUserGamesCookieValues(c)
	if err != nil {
		logger.Error("Error trying get cookie", err, zap.String("journey", "CreateComment Controller"))
		restErr := rest_err.NewUnauthorizedError("cookie inválido")
		c.JSON(restErr.Code, restErr)
		return
	}

	var createCommentRequest comment_request.CreateComment
	if err := c.ShouldBindJSON(&createCommentRequest); err != nil {
		logger.Error("Error trying convert fields", err, zap.String("journey", "CreateComment Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	commentDomain := comment_domain.NewCommentDomain("", createCommentRequest.VideoId, createCommentRequest.ParentId, createCommentRequest.AnswerId, cookie.Id, createCommentRequest.VideoComment, "")
	restErr := cc.commentService.CreateComment(commentDomain)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("Comment created with sucess!!!", zap.String("journey", "CreateComment Controller"))
	c.Status(http.StatusCreated)
}