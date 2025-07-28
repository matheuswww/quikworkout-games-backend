package judge_util_repository

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

func CheckJudgeToken(code, id, tokenType, journey string, mysql *sql.DB, ctx context.Context) *rest_err.RestErr {
	var judgeCode string
	var created_at string
	var retries int
	query := fmt.Sprintf("SELECT code,retries,created_at FROM judge_token WHERE token_type = '%s' AND judge_id = ?", tokenType)
	err := mysql.QueryRowContext(ctx, query, id).Scan(&judgeCode, &retries, &created_at)
	if err == sql.ErrNoRows {
		logger.Error("Error trying get code", err, zap.String("journey", journey))
		return rest_err.NewBadRequestError("você não possui um código registrado")
	} else if err != nil {
		logger.Error("Error trying scan queryRow", err, zap.String("journey", journey))
		return rest_err.NewInternalServerError("server error")
	}
	if retries >= 3 {
		logger.Error("Error max requests", errors.New("max requests"), zap.String("journey", journey))
		return rest_err.NewBadRequestError("maximo de tentativas atingido")
	}
	if judgeCode != code {
		logger.Error("Error expecting code", errors.New("wrong auth code"), zap.String("journey", journey))
		query = fmt.Sprintf("UPDATE judge_token SET retries = ? WHERE token_type = '%s' AND judge_id = ?", tokenType)
		retries++
		_, err := mysql.ExecContext(ctx, query, retries, id)
		if err != nil {
			logger.Error("Error trying update retries", err, zap.String("journey", journey))
			return rest_err.NewInternalServerError("server error")
		}
		return rest_err.NewBadRequestError("código inválido")
	}

	exp, err := checkExpTime(created_at, time.Minute*5)
	if err != nil {
		logger.Error("Error trying CheckExpTime", err, zap.String("journey", journey))
		return rest_err.NewInternalServerError("server error")
	}
	if exp {
		logger.Error("Error code was expirad", errors.New("code expirad"), zap.String("journey", journey))
		return rest_err.NewBadRequestError("código expirado")
	}
	return nil
}

func checkExpTime(created_at string, expTime time.Duration) (bool, error) {
	format := "2006-01-02 15:04:05"
	created_atFormated, err := time.Parse(format, created_at)
	if err != nil {
		return false, err
	}
	timeNowFormated, err := time.Parse(format, time.Now().Format(format))
	if err != nil {
		return false, err
	}
	exp := timeNowFormated.Sub(created_atFormated)
	if exp > expTime {
		return true, nil
	}
	return false, nil
}
