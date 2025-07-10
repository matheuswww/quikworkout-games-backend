package admin_repository

import (
	"context"

	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ar *adminRepository) HasPlacing(edition_id string) (bool, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var count int
	query := "SELECT COUNT(*) FROM participant WHERE edition_id = ?"
	err := ar.mysql.QueryRowContext(ctx, query, edition_id).Scan(&count)
	if err != nil {
		logger.Error("Error trying HasPlacing", err, zap.String("journey", "HasPlacing Repository"))
		return false, rest_err.NewInternalServerError("server error")
	}

	if count == 0 {
		return false, rest_err.NewBadRequestError("no participants were found")
	}

	count = 0
	query = "SELECT COUNT(*) FROM participant WHERE placing IS NOT NULL AND edition_id = ? LIMIT 1"
	err = ar.mysql.QueryRowContext(ctx, query, edition_id).Scan(&count)
	if err != nil {
		logger.Error("Error trying get participant", err, zap.String("journey", "HasPlacing Repository"))
		return false, rest_err.NewInternalServerError("server error")
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}