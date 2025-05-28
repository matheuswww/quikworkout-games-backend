package comment_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	comment_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/comment/request"
	"go.uber.org/zap"
)

func (cc *commentController) GetComment(c *gin.Context) {
	logger.Info("Init CreateComment controller", zap.String("journey", "GetComment Controller"))

	var getCommentRequest comment_request.GetComment
	if err := c.ShouldBind(&getCommentRequest); err != nil {
		logger.Error("Error trying convert fields", err, zap.String("journey", "GetComment Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	comments, restErr := cc.commentService.GetComment(getCommentRequest.VideoId, getCommentRequest.Cursor, getCommentRequest.CommentId)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, comments)
} 