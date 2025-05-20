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
		return rest_err.NewInternalServerError("não foi possivel salvar o pedido")
	}
	if sessionIdFromCookie != sessionId {
		logger.Error("Error invalid session", errors.New("invalid session"), zap.String("journey", "ChangePassword Repository"))
		return rest_err.NewUnauthorizedError("cookie inválido")
	}
	return nil
}
