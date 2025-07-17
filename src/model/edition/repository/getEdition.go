package edition_repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	edition_domain "github.com/matheuswww/quikworkout-games-backend/src/model/edition"
	"go.uber.org/zap"
)

func (er *editionRepository) GetEdition(number, limit int, cursor string) ([]edition_domain.EditionDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var args []any

	query := "SELECT e.edition_id, e.start_date, e.closing_date, e.rules, e.number, t.top, t.gain, t.category, c.challenge, c.category, c.sex, e.created_at FROM (SELECT edition_id, start_date, closing_date, rules, number, created_at FROM edition "
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
		query = query[:len(query)-4]
	}
	if limit > 10 || limit == 0 {
		limit = 10
	}
	query += "ORDER BY created_at DESC LIMIT ?) AS e LEFT JOIN challenge AS c ON c.edition_id = e.edition_id LEFT JOIN top AS t ON t.edition_id = e.edition_id AND t.category = c.category "
	args = append(args, limit)
	query += ""
	rows, err := er.mysql.QueryContext(ctx, query, args...)

	if err != nil {
		logger.Error("Error trying get edition", err, zap.String("journey", "GetEdition Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()

	var editionsDomain []edition_domain.EditionDomainInterface
	var editionDomain edition_domain.EditionDomainInterface
	var prevId string
	var tops []edition_domain.Top
	var challenges []edition_domain.Challenge
	challengeMap := make(map[string]bool)
	topMap := make(map[string]bool)
	
	for rows.Next() {
		var id, start_date, closing_date, rules, challenge, challengeCategory, sex, created_at string
		var number int
		var topCategory sql.NullString
		var gain, top sql.NullInt64
		err := rows.Scan(&id, &start_date, &closing_date, &rules, &number, &top, &gain, &topCategory, &challenge, &challengeCategory, &sex, &created_at)
		if err != nil {
			logger.Error("Error trying scan row", err, zap.String("journey", "GetEdition Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		if prevId != id {
			if editionDomain != nil {
				editionDomain.SetTops(tops)
				editionDomain.SetChallenge(challenges)
				editionsDomain = append(editionsDomain, editionDomain)
			}
			editionDomain = edition_domain.NewEditionDomain(id, start_date, closing_date, rules, nil, nil, number, created_at)
			tops = nil
		}
		if _, exists := challengeMap[challengeCategory+challenge+sex]; !exists {
			challenges = append(challenges, edition_domain.Challenge{
				Sex: sex,
				Challenge: challenge,
				Category:  challengeCategory,
			})
			challengeMap[challengeCategory+challenge+sex] = true
		}
		if top.Valid && topCategory.Valid && gain.Valid {
			topKey := fmt.Sprintf("%s%d", topCategory.String, top.Int64)
			if _, exists := topMap[topKey]; !exists {
				tops = append(tops, edition_domain.Top{
						Top:  int(top.Int64),
						Gain: int(gain.Int64),
						Category: topCategory.String,
					})
					topMap[topKey] = true
				}
			}
			prevId = id
		}
	if editionDomain != nil {
		editionDomain.SetTops(tops)
		editionDomain.SetChallenge(challenges)
		editionsDomain = append(editionsDomain, editionDomain)
	}
	if len(editionsDomain) == 0 {
		return nil, rest_err.NewNotFoundError("no edition found")
	}
	return editionsDomain, nil
}
