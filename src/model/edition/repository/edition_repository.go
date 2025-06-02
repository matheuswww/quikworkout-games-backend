package edition_repository

import (
	"database/sql"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	edition_domain "github.com/matheuswww/quikworkout-games-backend/src/model/edition"
)

type editionRepository struct {
	mysql *sql.DB
}

func NewEditionRepository(mysql *sql.DB) EditionRepository {
	return &editionRepository{
		mysql,
	}
}

type EditionRepository interface {
	GetEdition(number, limit int, cursor string) ([]edition_domain.EditionDomainInterface, *rest_err.RestErr)
}
