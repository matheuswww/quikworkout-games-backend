package judge_repository

import (
	"database/sql"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	judge_domain "github.com/matheuswww/quikworkout-games-backend/src/model/judge"
)

type judgeRepository struct {
	mysql *sql.DB
}

func NewJudgeRepository(mysql *sql.DB) JudgeRepository {
	return &judgeRepository{
		mysql,
	}
}

type JudgeRepository interface {
	SendSigninCode(judge judge_domain.JudgeDomainInterface, code string) *rest_err.RestErr
	CheckSigninCode(judgeDomain judge_domain.JudgeDomainInterface, code string) *rest_err.RestErr
}
