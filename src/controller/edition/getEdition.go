package edition_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	default_validator "github.com/matheuswww/quikworkout-games-backend/src/configuration/validation/defaultValidator"
	edition_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/edition/request"
	edition_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/edition/response"
	"go.uber.org/zap"
)

func (ec *editionController) GetEdition(c *gin.Context) {
	logger.Info("Init GetEdition", zap.String("journey", "GetEdition Controller"))
	var getEditionRequest edition_request.GetEdition
	if err := c.ShouldBindQuery(&getEditionRequest); err != nil {
		logger.Error("Error trying convert fileds", err, zap.String("journey", "GetEdition Controller"))
		restErr := default_validator.HandleDefaultValidatorErrors(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	editions, restErr := ec.editionService.GetEdition(getEditionRequest.Number, getEditionRequest.Limit, getEditionRequest.Cursor)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	var editionsJson []edition_response.Edition
	for _,edition := range editions {
		var tops []edition_response.Top
		for _,top := range edition.GetTops() {
			tops = append(tops, edition_response.Top{
				Gain: top.Gain,
				Top: top.Top,		
			})
		}
		editionsJson = append(editionsJson, edition_response.Edition{
			Id: edition.GetId(),
			StartDate: edition.GetStartDate(),
			ClosingDate: edition.GetClosingDate(),
			Rules: edition.GetRules(),
			Challenge: edition.GetChallenge(),
			Number: edition.GetNumber(),
			Tops: tops,
			CreatedAt: edition.GetCreatedAt(),
		})
	}
	
	c.JSON(http.StatusOK, editionsJson)
}