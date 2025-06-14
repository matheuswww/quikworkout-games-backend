package admin_repository

import (
	"context"

	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ar *adminRepository) HasNoPlacing(edition_id string) (bool, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var count int
	query := "SELECT COUNT(*) FROM participant WHERE edition_id = ?"
	err := ar.mysql.QueryRowContext(ctx, query, edition_id).Scan(&count)
	if err != nil {
		logger.Error("Error trying HasNoPlacing", err, zap.String("journey", "HasNoPlacing Repository"))
		return false, rest_err.NewInternalServerError("server error")
	}

	if count == 0 {
		return false, rest_err.NewBadRequestError("no participants were found")
	}

	count = 0
	query = "SELECT COUNT(*) FROM participant WHERE desqualified IS NULL AND placing IS NULL AND edition_id = ?"
	err = ar.mysql.QueryRowContext(ctx, query, edition_id).Scan(&count)
	if err != nil {
		logger.Error("Error trying HasNoPlacing", err, zap.String("journey", "HasNoPlacing Repository"))
		return false, rest_err.NewInternalServerError("server error")
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}