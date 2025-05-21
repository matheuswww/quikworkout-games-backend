package edition_controller

import (
	"github.com/gin-gonic/gin"
	edition_service "github.com/matheuswww/quikworkout-games-backend/src/model/edition/service"
)

type editionController struct {
	editionService edition_service.EditionService
}

func NewEditionController(editionService edition_service.EditionService) EditionController {
	return &editionController{
		editionService,
	}
}

type EditionController interface {
	GetEdition(c *gin.Context)
}
