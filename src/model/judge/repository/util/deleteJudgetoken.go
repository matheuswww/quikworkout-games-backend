package judge_util_repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

func DeleteJudgeToken(id, tokenType, journey string, mysql *sql.DB, ctx context.Context) *rest_err.RestErr {
	query := fmt.Sprintf("DELETE FROM judge_token WHERE token_type = '%s' AND judge_id = ?", tokenType)
	_, err := mysql.ExecContext(ctx, query, id)
	if err != nil {
		logger.Info("Error trying delete judge_token", zap.String("journey", journey))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}
