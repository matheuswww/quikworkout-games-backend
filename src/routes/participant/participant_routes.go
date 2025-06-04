package participant_router

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	participant_controller "github.com/matheuswww/quikworkout-games-backend/src/controller/participant"
	user_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user"
	user_games_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user/user_games"
	participant_repository "github.com/matheuswww/quikworkout-games-backend/src/model/participant/repository"
	participant_service "github.com/matheuswww/quikworkout-games-backend/src/model/participant/service"
	"go.uber.org/zap"
)

func InitParticipantRoutes(r *gin.RouterGroup, database *sql.DB) {
	participantController := initParticipantRoutes(database)
	cookieStore, err := user_cookie.Store()
	if err != nil {
		logger.Error("Error loading cookie store", err, zap.String("journey", "InitUserRoutes"))
		log.Fatal("Error cookie store")
	}
	sessionNames := []string{user_games_cookie.SessionUserGames}
	r.Use(sessions.SessionsMany(sessionNames, cookieStore))
	r.POST("/participant/createParticipant", participantController.CreateParticipant)
	r.GET("/participant/getParticipants", participantController.GetParticipant)
	r.GET("/participant/hasTicket", participantController.HasTicket)
	r.POST("/participant/videoSent", participantController.VideoSent)
}

func initParticipantRoutes(database *sql.DB) participant_controller.ParticipantController {
	participantRepository := participant_repository.NewParticipantRepository(database)
	participantService := participant_service.NewParticipantService(participantRepository)
	participantController := participant_controller.NewParticipantController(participantService)
	return participantController
}
