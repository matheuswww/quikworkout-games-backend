package judge_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	judge_domain "github.com/matheuswww/quikworkout-games-backend/src/model/judge"
	judge_repository "github.com/matheuswww/quikworkout-games-backend/src/model/judge/repository"
)

type judgeService struct {
	judgeRepository judge_repository.JudgeRepository
}

func NewJudgeService(judgeRepository judge_repository.JudgeRepository) JudgeService {
	return &judgeService{
		judgeRepository,
	}
}

type JudgeService interface {
	SendSigninCode(judgeModel judge_domain.JudgeDomainInterface, token string) *rest_err.RestErr 
	CheckSigninCode(judge judge_domain.JudgeDomainInterface, code, token string) *rest_err.RestErr
}
