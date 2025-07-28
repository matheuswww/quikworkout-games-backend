package judge_repository

import (
	"context"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	judge_domain "github.com/matheuswww/quikworkout-games-backend/src/model/judge"
	judge_util_repository "github.com/matheuswww/quikworkout-games-backend/src/model/judge/repository/util"
	"go.uber.org/zap"
)

func (ar *judgeRepository) CheckSigninCode(judgeDomain judge_domain.JudgeDomainInterface, code string) *rest_err.RestErr {
	journey := "CheckSigninCode Repository"
	logger.Info("Init CheckSigninCode Repository", zap.String("journey", journey))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	tokenType := "signin"
	err := judge_util_repository.CheckJudgeToken(code, judgeDomain.GetId(), tokenType, journey, ar.mysql, ctx)
	if err != nil {
		return err
	}
	err = judge_util_repository.DeleteJudgeToken(judgeDomain.GetId(), tokenType, journey, ar.mysql, ctx)
	if err != nil {
		return err
	}
	return nil
}
