package judge_util_repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

func InsertJudgeToken(code, id, tokenType, journey string, mysql *sql.DB, ctx context.Context) *rest_err.RestErr {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) judge_id FROM judge_token WHERE token_type = '%s' AND judge_id = ?", tokenType)
	err := mysql.QueryRowContext(ctx, query, id).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error trying scan queryRow", err, zap.String("journey", journey))
		return rest_err.NewInternalServerError("server error")
	}
	if count > 0 {
		format := "2006-01-02 15:04:05"
		created_at, err := time.Parse(format, time.Now().Format(format))
		if err != nil {
			logger.Error("Error trying generate timestamp", err, zap.String("journey", journey))
			return rest_err.NewInternalServerError("server error")
		}
		query := fmt.Sprintf("UPDATE judge_token SET code = ?,retries = 0, created_at = ? WHERE token_type = '%s' AND judge_id = ?", tokenType)
		_, err = mysql.ExecContext(ctx, query, code, created_at, id)
		if err != nil {
			logger.Error("Error trying update token_token", err, zap.String("journey", journey))
			return rest_err.NewInternalServerError("server error")
		}
		return nil
	}
	query = "INSERT INTO judge_token (judge_id, code, token_type) VALUES (?, ?, ?)"
	_, err = mysql.ExecContext(ctx, query, id, code, tokenType)
	if err != nil {
		logger.Error("Error trying insert token_type in mysql", err, zap.String("journey", journey))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}
