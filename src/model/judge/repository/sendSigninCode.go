package judge_repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	judge_domain "github.com/matheuswww/quikworkout-games-backend/src/model/judge"
	judge_util_repository "github.com/matheuswww/quikworkout-games-backend/src/model/judge/repository/util"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (jr *judgeRepository) SendSigninCode(judge judge_domain.JudgeDomainInterface, code string) *rest_err.RestErr {
	journey := "SendSigninCode Repository"
	logger.Info("Init SendSigninCode Repository", zap.String("journey", journey))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	query := "SELECT judge_id,name,password FROM judge WHERE email = ?"
	var id, name string
	var encryptedPassword []byte
	err := jr.mysql.QueryRowContext(ctx, query, judge.GetEmail()).Scan(&id, &name, &encryptedPassword)
	if err == sql.ErrNoRows {
		logger.Error("Error trying get judge, email not found", err, zap.String("journey", journey))
		return rest_err.NewBadRequestError("email ou senha incorretos")
	} else if err != nil {
		logger.Error("Error trying scan queryRow", err, zap.String("journey", journey))
		return rest_err.NewInternalServerError("server error")
	}
	err = bcrypt.CompareHashAndPassword(encryptedPassword, []byte(judge.GetPassword()))
	if err != nil {
		logger.Error("Error trying signin judge wrong password", err, zap.String("journey", journey))
		return rest_err.NewBadRequestError("email ou senha incorretos")
	}
	judge.SetName(name)
	judge.SetId(id)
	tokenType := "signin"
	restErr := judge_util_repository.InsertJudgeToken(code, judge.GetId(), tokenType, journey, jr.mysql, ctx)
	if restErr != nil {
		return restErr
	}
	return nil
}
