package admin_repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)
func (ar *adminRepository) GrantTicket(user string) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := "SELECT user_id FROM user_games WHERE user = ?"
	var user_id string
	err := ar.mysql.QueryRowContext(ctx, query, user).Scan(&user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error trying get user_games", err, zap.String("journey", "GrantTicket Repository"))
			return rest_err.NewBadRequestError("user not found")
		}
		logger.Error("Error trying get user_games", err, zap.String("journey", "GrantTicket Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	query = "SELECT edition_id FROM edition ORDER BY created_at DESC LIMIT 1"
	var edition_id string
	err = ar.mysql.QueryRowContext(ctx, query).Scan(&edition_id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error trying get edition", err, zap.String("journey", "GrantTicket Repository"))
			return rest_err.NewBadRequestError("no edition found")
		}
		logger.Error("Error trying get edition", err, zap.String("journey", "GrantTicket Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	query = "SELECT 1 FROM direct_ticket WHERE user_id = ? AND edition_id = ?"
	var exists int
	err = ar.mysql.QueryRowContext(ctx, query, user_id, edition_id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error trying get direct tikect", err, zap.String("journey", "GrantTicket Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if exists == 1 {
		logger.Error("User already has a direct ticket", nil, zap.String("journey", "GrantTicket Repository"))
		return rest_err.NewBadRequestError("user already has a direct ticket")
	}

	query = "INSERT INTO direct_ticket (user_id, edition_id) VALUES (?, ?)"
	_, err = ar.mysql.ExecContext(ctx, query, user_id, edition_id)
	if err != nil {
		logger.Error("Error trying insert direct ticket", err, zap.String("journey", "GrantTicket Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	
	return nil
}