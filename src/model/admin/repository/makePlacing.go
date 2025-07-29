package admin_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

type top struct {
	top int
	gain int	
}

func (ar *adminRepository) MakePlacing(editionId, category, sex string) *rest_err.RestErr {
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
	query = "SELECT COUNT(*) FROM participant WHERE edition_id = ? AND category = ? AND sex = ? AND desqualified IS NULL AND checked IS FALSE"
	err = ar.mysql.QueryRowContext(ctx, query, editionId, category, sex).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return rest_err.NewNotFoundError("no participants found")
		}
		logger.Error("Error trying get participants", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if count > 0 {
		logger.Error("Error trying MakePlacing", errors.New("there are still participants without a time"), zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewBadRequestError("there are still participants without checked")
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
	
	tx, err := ar.mysql.Begin()
	if err != nil {
		logger.Error("Error trying Begin", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	
	restErr := updateFinalTime(ctx, tx, editionId, category, sex)
	if restErr != nil {
		return restErr
	}

	userIds, restErr := updatePlacing(ctx, tx, editionId, category, sex)
	if restErr != nil {
		return  restErr
	}

	restErr = updateEarnings(ctx, tx, userIds, tops)
	if restErr != nil {
		return  restErr
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Error trying Commit", err, zap.String("journey", "MakePlacing Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	return nil
}

func updateFinalTime(ctx context.Context, tx *sql.Tx, editionId, category, sex string) (*rest_err.RestErr) {
	query := `SELECT user_id, user_time FROM participant WHERE edition_id = ? AND category = ? AND sex = ? AND desqualified IS NULL AND final_time IS NULL ORDER BY user_time ASC`
	rows, err := tx.QueryContext(ctx, query, editionId, category, sex)
	if err != nil {
		logger.Error("Error trying get participants", err, zap.String("journey", "MakePlacing Repository"))
		_ = tx.Rollback()
		return rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()

	var userIds []string
	var times []string
	for rows.Next() {
		var userId, time string
		if err := rows.Scan(&userId, &time); err != nil {
			logger.Error("Error trying Scan", err, zap.String("journey", "MakePlacing Repository"))
			_ = tx.Rollback()
			return rest_err.NewInternalServerError("server error")
		}
		userIds = append(userIds, userId)
		times = append(times, time)
	}

	if(len(userIds) == 0) {
		return nil
	}

	query = "UPDATE participant SET final_time = CASE user_id "
	endQuery := "END WHERE edition_id = ? AND category = ? AND sex = ? AND user_id IN ("
	args := []any{}
	endArgs := []any{editionId, category, sex}

	for i, userId := range userIds {
		query += "WHEN ? THEN ? "
		args = append(args, userId, times[i])
		if i > 0 {
			endQuery += ", "
		}
		endQuery += "?"
		endArgs = append(endArgs, userId)
	}
	endQuery += ")"
	query += endQuery
	args = append(args, endArgs...)
	fmt.Println(query)
	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		logger.Error("Error updating final_time", err, zap.String("journey", "MakePlacing Repository"))
		_ = tx.Rollback()
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}

func updatePlacing(ctx context.Context, tx *sql.Tx, editionId, category, sex string) ([]string, *rest_err.RestErr) {
	query := `SELECT user_id FROM participant WHERE edition_id = ? AND category = ? AND sex = ? AND desqualified IS NULL ORDER BY final_time ASC`
	rows, err := tx.QueryContext(ctx, query, editionId, category, sex)
	if err != nil {
		logger.Error("Error trying get participants", err, zap.String("journey", "MakePlacing Repository"))
		_ = tx.Rollback()
		return nil, rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()

	var userIds []string
	for rows.Next() {
		var userId string
		if err := rows.Scan(&userId); err != nil {
			logger.Error("Error trying Scan", err, zap.String("journey", "MakePlacing Repository"))
			_ = tx.Rollback()
			return nil, rest_err.NewInternalServerError("server error")
		}
		userIds = append(userIds, userId)
	}


	query = "UPDATE participant SET placing = CASE user_id "
	endQuery := "END WHERE edition_id = ? AND category = ? AND sex = ? AND user_id IN ("
	args := []any{}
	endArgs := []any{editionId, category, sex}

	for i, userId := range userIds {
		query += "WHEN ? THEN ? "
		args = append(args, userId, i+1)
		if i > 0 {
			endQuery += ", "
		}
		endQuery += "?"
		endArgs = append(endArgs, userId)
	}
	endQuery += ")"
	query += endQuery
	args = append(args, endArgs...)

	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		logger.Error("Error updating placing", err, zap.String("journey", "MakePlacing Repository"))
		_ = tx.Rollback()
		return userIds, rest_err.NewInternalServerError("server error")
	}
	return userIds, nil
}

func updateEarnings (ctx context.Context, tx *sql.Tx, userIds []string, tops []top,) *rest_err.RestErr {
	if len(tops) == 0 {
		return nil
	}
	query := "UPDATE user_games SET earnings = CASE user_id "
	endQuery := "END WHERE user_id IN ("
	args := []any{}
	endArgs := []any{}

	for i, userId := range userIds {
		if i >= len(tops) {
			break
		}
		query += "WHEN ? THEN earnings + ? "
		args = append(args, userId, tops[i].gain)
		if i > 0 {
			endQuery += ", "
		}
		endQuery += "?"
		endArgs = append(endArgs, userId)
	}
	endQuery += ")"
	query += endQuery
	args = append(args, endArgs...)

	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		logger.Error("Error updating earnings", err, zap.String("journey", "MakePlacing Repository"))
		_ = tx.Rollback()
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}