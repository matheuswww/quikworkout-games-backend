package judge_service

import (
	"errors"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/recaptcha"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	judge_domain "github.com/matheuswww/quikworkout-games-backend/src/model/judge"
	"go.uber.org/zap"
)

func (js *judgeService) CheckSigninCode(judge judge_domain.JudgeDomainInterface, code, token string) *rest_err.RestErr {
	logger.Info("Init CheckSigninCode Service", zap.String("journey", "CheckSigninCode Service"))
	recaptchaErr := recaptcha.NewRecaptcha().ValidateRecaptcha(token)
	if recaptchaErr != nil {
		return recaptchaErr
	}
	err := js.judgeRepository.CheckSigninCode(judge, code)
	if err != nil {
		logger.Error("Error trying CheckSigninCode Repository", errors.New("CheckSigninCode Repository Error"), zap.String("journey", "CheckSigninCode Service"))
		return err
	}
	return nil
}
