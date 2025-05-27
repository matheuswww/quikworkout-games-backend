package comment_controller

import (
	"github.com/gin-gonic/gin"
	comment_service "github.com/matheuswww/quikworkout-games-backend/src/model/comment/service"
)

type commentController struct {
	commentService comment_service.CommentService
}

type CommentController interface {
	CreateComment(c *gin.Context)
}

func NewCommentController(commentService comment_service.CommentService) CommentController {
	return &commentController{
		commentService,
	}
}
