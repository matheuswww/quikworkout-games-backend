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

type top struct {
	top int
	gain int	
}

func (ar *adminRepository) MakePlacing(editionId, category string) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var closing_date string
	query := "SELECT closing_date FROM edition WHERE edition_id = ?"
	err := ar.mysql.QueryRowContext(ctx, query, editionId).Scan(&closing_date)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error trying get edition", errors.New("edition not found"), zap.String("journey", "MakePlacing repository"))
			return rest_err.NewBadRequestError("edition not found")
		}
		logger.Error("Error trying get edition", err, zap.String("journey", "MakePlacing repository"))
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
	query = "SELECT COUNT(*) FROM participant WHERE edition_id = ? AND category = ? AND desqualified IS NULL AND (user_time IS NULL OR checked IS FALSE)"
	err = ar.mysql.QueryRowContext(ctx, query, editionId, category).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return rest_err.NewNotFoundError("no participants found")
		}
		logger.Error("Error trying get participants", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if count > 0 {
		logger.Error("Error trying MakePlacing", errors.New("there are still participants without a time"), zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewBadRequestError("there are still participants without a time or checked")
	}

	query = "SELECT top, gain FROM top WHERE edition_id = ? AND category = ? ORDER BY top ASC"
	rows, err := ar.mysql.QueryContext(ctx, query, editionId, category)
	if err != nil {
		logger.Error("Error trying get tops", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()
	
	var tops []top
	for rows.Next() {
		var tp, gain int
		err := rows.Scan(&tp, &gain)
		if err != nil {
			logger.Error("Error trying rows next", err, zap.String("journey", "MakePlacing Repository"))
			return rest_err.NewInternalServerError("server error")
		}
		tops = append(tops, top{
			top: tp,
			gain: gain,
		})
	}
	
	query = "SELECT user_id FROM participant WHERE edition_id = ? AND category = ? AND desqualified IS NULL ORDER BY user_time ASC"
	rows, err = ar.mysql.QueryContext(ctx, query, editionId, category)
	if err != nil {
		logger.Error("Error trying get participants", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()
	
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

	tx, err := ar.mysql.Begin()
	if err != nil {
		logger.Error("Error trying Begin", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	query = "UPDATE participant SET placing = CASE user_id "
	endQuery := "END WHERE edition_id = ? AND category = ? AND user_id IN ("
	endQueryArgs := []any{editionId, category}
	args := []any{}

	for i, userId := range userIds {
		query += "WHEN ? THEN ? "
		if i > 0 {
			endQuery += ", "
		}
		endQuery += "?"
		endQueryArgs = append(endQueryArgs, userId)
		args = append(args, userId, i+1)
	}
	endQuery += ")"
	query += endQuery
	args = append(args, endQueryArgs...)

	_,err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying update participants", err, zap.String("journey", "MakePlacing Repository"))
		err := tx.Rollback()
		if err != nil {
			logger.Error("Error trying rollback", err, zap.String("journey", "MakePlacing Repository"))
			return rest_err.NewInternalServerError("server error")
		}
		return rest_err.NewInternalServerError("server error")
	}
	args = nil
	endQueryArgs = nil

	if(len(tops) > 0) {
		query = "UPDATE user_games SET earnings = CASE user_id "
		endQuery = "END WHERE user_id IN ("
		for i, userId := range userIds {
			if i <= len(tops) - 1 {
			query += "WHEN ? THEN earnings + ? "
			args = append(args, userId, tops[i].gain)
			if i > 0 {
				endQuery += ", "
			}
			endQuery += "?"
			endQueryArgs = append(endQueryArgs, userId)
			continue
		}
		break
	}
	endQuery += ")"
	query += endQuery

	args = append(args, endQueryArgs...)
	_,err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying update participants", err, zap.String("journey", "MakePlacing Repository"))
		err := tx.Rollback()
		if err != nil {
			logger.Error("Error trying rollback", err, zap.String("journey", "MakePlacing Repository"))
			return rest_err.NewInternalServerError("server error")
		}
		return rest_err.NewInternalServerError("server error")
	}
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Error trying Commit", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	return nil
}