package edition_router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	edition_controller "github.com/matheuswww/quikworkout-games-backend/src/controller/edition"
	edition_repository "github.com/matheuswww/quikworkout-games-backend/src/model/edition/repository"
	edition_service "github.com/matheuswww/quikworkout-games-backend/src/model/edition/service"
)

func InitEditionRoutes(r *gin.RouterGroup, database *sql.DB) {
	editionController := initEditionRoutes(database)

	r.GET("/edition/getEdition", editionController.GetEdition)
	
}

func initEditionRoutes(database *sql.DB) edition_controller.EditionController {
	editionRepository := edition_repository.NewEditionRepository(database)
	editionService := edition_service.NewEditionService(editionRepository)
	editionController := edition_controller.NewEditionController(editionService)
	return editionController
}
