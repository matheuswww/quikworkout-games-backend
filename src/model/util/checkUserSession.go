package model_util

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

var db *sql.DB

func InitDb(database *sql.DB) {
	db = database
}

func CheckUser(sessionIdFromCookie, user_id string) *rest_err.RestErr {
	logger.Info("Init CheckUser", zap.String("journey", "CheckUser"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var sessionId string
	query := "SELECT session_id FROM user WHERE user_id = ?"
	err := db.QueryRowContext(ctx, query, user_id).Scan(&sessionId)
	if err != nil {
		logger.Error("Error trying CheckUser", err, zap.String("journey", "CheckUser"))
		return rest_err.NewInternalServerError("server error")
	}
	if sessionIdFromCookie != sessionId {
		logger.Error("Error invalid session", errors.New("invalid session"), zap.String("journey", "CheckUser"))
		return rest_err.NewUnauthorizedError("cookie inválido")
	}
	return nil
}

func CheckUserGames(sessionIdFromCookie, user_id string) *rest_err.RestErr {
	logger.Info("Init CheckUser", zap.String("journey", "CheckUser"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var sessionId string
	query := "SELECT session_id FROM user_games WHERE user_id = ?"
	err := db.QueryRowContext(ctx, query, user_id).Scan(&sessionId)
	if err != nil {
		logger.Error("Error trying CheckUser", err, zap.String("journey", "CheckUser"))
		return rest_err.NewInternalServerError("server error")
	}
	if sessionIdFromCookie != sessionId {
		logger.Error("Error invalid session", errors.New("invalid session"), zap.String("journey", "CheckUser Repository"))
		return rest_err.NewUnauthorizedError("cookie inválido")
	}
	return nil
}
