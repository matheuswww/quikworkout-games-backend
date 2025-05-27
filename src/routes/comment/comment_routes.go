package comment_router

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	comment_controller "github.com/matheuswww/quikworkout-games-backend/src/controller/comment"
	user_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	comment_repository "github.com/matheuswww/quikworkout-games-backend/src/model/comment/repository"
	comment_service "github.com/matheuswww/quikworkout-games-backend/src/model/comment/service"
	"go.uber.org/zap"
)

func InitCommentRoutes(r *gin.RouterGroup, database *sql.DB) {
	commentController := initCommentRoutes(database)
	cookieStore, err := user_cookie.Store()
	if err != nil {
		logger.Error("Error loading cookie store", err, zap.String("journey", "InitCommentRoutes"))
		log.Fatal("Error cookie store")
	}
	sessionNames := []string{user_games_cookie.SessionUserGames}
	r.Use(sessions.SessionsMany(sessionNames, cookieStore))

	r.POST("/comment/createComment", commentController.CreateComment)
}

func initCommentRoutes(database *sql.DB) comment_controller.CommentController {
	commentRepository := comment_repository.NewCommentRepository(database)
	commentService := comment_service.NewCommentService(commentRepository)
	commentController := comment_controller.NewCommentController(commentService)
	return commentController
}
