package admin_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)


func (ar *adminRepository) MakePlacing(editionId string) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var closing_date string
	query := "SELECT closing_date FROM edition WHERE edition_id = ?"
	err := ar.mysql.QueryRowContext(ctx, query, editionId).Scan(&closing_date)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error Edition not found", errors.New("edition not found"), zap.String("journey", "MakePlacing repository"))
			return rest_err.NewNotFoundError("edition not found")
		}
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "MakePlacing repository"))
		return rest_err.NewInternalServerError("server error")
	}
	format := "2006-01-02"
	closing_date_formated, err := time.Parse(format, closing_date)
	if err != nil {
		logger.Error("Error trying Parse date", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	closing_date_formated = closing_date_formated.Add(24*time.Hour - time.Second)
	now := time.Now()
	if now.Before(closing_date_formated) {
		logger.Error("Error trying MakePlacing", errors.New("the edition has not yet been closed"), zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewBadRequestError("the edition has not yet been closed")
	}

	var count int
	query = "SELECT COUNT(*) user_time FROM participant WHERE edition_id = ? AND user_time IS NULL AND desqualified IS NULL"
	err = ar.mysql.QueryRowContext(ctx, query, editionId).Scan(&count)
	if err != nil {
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if count > 0 {
		logger.Error("Error trying MakePlacing", errors.New("there are still participants without a time"), zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewBadRequestError("there are still participants without a time")
	}
 
	query = "SELECT user_id FROM participant WHERE edition_id = ? AND desqualified IS NULL ORDER BY user_time ASC"
	rows, err := ar.mysql.QueryContext(ctx, query, editionId)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error trying QueryContext", errors.New("no participants found"), zap.String("journey", "MakePlacing Repository"))
			return rest_err.NewNotFoundError("no participants found")
		}
		logger.Error("Error trying QueryContext", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	var userIds []string
	for rows.Next() {
		var user_id string
		err := rows.Scan(&user_id)
		if err != nil {
			logger.Error("Error trying Scan", err, zap.String("journey", "MakePlacing Repository"))
			return rest_err.NewInternalServerError("server error")
		}
		userIds = append(userIds, user_id)
	}
	
	query = "UPDATE participant SET placing = CASE user_id "
	args := []any{}

	for i, userId := range userIds {
		query += "WHEN ? THEN ? "
		args = append(args, userId, i+1)
	}

	query += "END WHERE edition_id = ? AND user_id IN ("

	args = append(args, editionId)

	for i, userID := range userIds {
		if i > 0 {
			query += ", "
		}
		query += "?"
		args = append(args, userID)
	}
	query += ")"

	_,err = ar.mysql.ExecContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying Scan", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	return nil
}