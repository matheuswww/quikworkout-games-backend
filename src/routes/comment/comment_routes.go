package comment_router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	comment_controller "github.com/matheuswww/quikworkout-games-backend/src/controller/comment"
	comment_repository "github.com/matheuswww/quikworkout-games-backend/src/model/comment/repository"
	comment_service "github.com/matheuswww/quikworkout-games-backend/src/model/comment/service"
)

func InitCommentRoutes(r *gin.RouterGroup, database *sql.DB) {
	_ = initCommentRoutes(database)
}

func initCommentRoutes(database *sql.DB) comment_controller.CommentController {
	commentRepository := comment_repository.NewCommentRepository(database)
	commentService := comment_service.NewCommentService(commentRepository)
	commentController := comment_controller.NewCommentController(commentService)
	return commentController
}
