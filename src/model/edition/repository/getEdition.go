package edition_repository

import (
	"context"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	edition_domain "github.com/matheuswww/quikworkout-games-backend/src/model/edition"
	"go.uber.org/zap"
)

func (er *editionRepository) GetEdition(number int, cursor string) ([]edition_domain.EditionDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var args []any

	query := "SELECT e.edition_id, e.start_date, e.closing_date, e.rules, e.challenge, e.number, t.top, t.gain, e.created_at FROM (SELECT edition_id, start_date, closing_date, rules, challenge, number, created_at FROM edition "
	if number != 0 || cursor != "" {
		query += "WHERE "
	}
	if number != 0 {
		query += "number = ? AND "
		args = append(args, number)
	}

	if cursor != "" {
		query += "created_at < ? AND "
		args = append(args, cursor)
	}
	if len(args) > 0 {
		query = query[:len(query) - 4]
	}
	query += "ORDER BY created_at DESC LIMIT 10) "
	query += "AS e JOIN top AS t ON t.edition_id = e.edition_id"

	rows, err := er.mysql.QueryContext(ctx, query, args...)

	if err != nil {
		logger.Error("Error trying GetEdition Repository", err, zap.String("journey", "GetEdition Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()
	
	var editionDomain []edition_domain.EditionDomainInterface
	var edition edition_domain.EditionDomainInterface
	var tops []edition_domain.Top
	var prevId string

	for rows.Next() {
		var id, start_date, closing_date, rules, challenge, created_at string
		var gain, top, number int
		err := rows.Scan(&id, &start_date, &closing_date, &rules, &challenge, &number, &gain, &top, &created_at)
		if err != nil {
			logger.Error("Error trying scan row", err, zap.String("journey", "GetEdition Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		if prevId != id {
			if edition != nil {
				edition.SetTops(tops)
				editionDomain = append(editionDomain, edition)
			}
			edition = edition_domain.NewEditionDomain(id, start_date, closing_date, rules, challenge, nil, number, created_at)
			tops = nil
		}
		tops = append(tops, edition_domain.Top{
			Top: top,
			Gain: gain,
		})
		prevId = id
	}
	if edition != nil {
		edition.SetTops(tops)
		editionDomain = append(editionDomain, edition)
	}
	if len(editionDomain) == 0 {
		return nil, rest_err.NewBadRequestError("no edition found")
	}
	return editionDomain, nil
}


