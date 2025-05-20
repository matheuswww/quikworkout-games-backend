package user_router

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	user_controller "github.com/matheuswww/quikworkout-games-backend/src/controller/user"
	user_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	user_proflie_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_profile"
	user_repository "github.com/matheuswww/quikworkout-games-backend/src/model/user/repository"
	user_service "github.com/matheuswww/quikworkout-games-backend/src/model/user/service"
	"go.uber.org/zap"
)

func InitUserRoutes(r *gin.RouterGroup, database *sql.DB) {
	userController := initUserRoutes(database)
	cookieStore, err := user_cookie.Store()
	if err != nil {
		logger.Error("Error loading cookie store", err, zap.String("journey", "InitUserTwoAuthRoutes"))
		log.Fatal("Error cookie store")
	}
	sessionNames := []string{user_proflie_cookie.SessionUserProfile, user_games_cookie.SessionUserGames}
	r.Use(sessions.SessionsMany(sessionNames, cookieStore))
	r.POST("/account/createAccount", userController.CreateAccount)
}

func initUserRoutes(database *sql.DB) user_controller.UserController {
	userRepository := user_repository.NewUserRepository(database)
	userService := user_service.NewUserService(userRepository)
	userController := user_controller.NewUserController(userService)
	return userController
}
