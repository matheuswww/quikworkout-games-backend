package edition_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	edition_domain "github.com/matheuswww/quikworkout-games-backend/src/model/edition"
	edition_repository "github.com/matheuswww/quikworkout-games-backend/src/model/edition/repository"
)

type editionService struct {
	editionRepository edition_repository.EditionRepository
}

func NewEditionService(editionRepository edition_repository.EditionRepository) EditionService {
	return &editionService{
		editionRepository,
	}
}

type EditionService interface {
	GetEdition(number, limit int, cursor string) ([]edition_domain.EditionDomainInterface, *rest_err.RestErr)
}
