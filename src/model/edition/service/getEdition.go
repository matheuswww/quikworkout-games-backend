package edition_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	edition_domain "github.com/matheuswww/quikworkout-games-backend/src/model/edition"
)

func (us *editionService) GetEdition(number int, cursor string) ([]edition_domain.EditionDomainInterface, *rest_err.RestErr) {
	return us.editionRepository.GetEdition(number, cursor)
}