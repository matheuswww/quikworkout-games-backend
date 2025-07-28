package judge_service

import (
	"fmt"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/recaptcha"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	judge_domain "github.com/matheuswww/quikworkout-games-backend/src/model/judge"
	judge_util_service "github.com/matheuswww/quikworkout-games-backend/src/model/judge/service/util"
	"go.uber.org/zap"
)

var (
	twoAuth = "contact@quikworkout.com.br"
)

func (js *judgeService) SendSigninCode(judgeModel judge_domain.JudgeDomainInterface, token string) *rest_err.RestErr {
	journey := "SendSigninCode Service"
	logger.Info("Init SendSigninCode Service", zap.String("journey", journey))
	recaptchaErr := recaptcha.NewRecaptcha().ValidateRecaptcha(token)
	if recaptchaErr != nil {
		return recaptchaErr
	}
	code, err := judge_util_service.GenerateCode()
	if err != nil {
		logger.Error("Error trying GenerateCode", err, zap.String("journey", journey))
		return rest_err.NewInternalServerError("server error")
	}
	restErr := js.judgeRepository.SendSigninCode(judgeModel, code)
	if restErr != nil {
		return restErr
	}
	title := "Codigo para login"
	msg := fmt.Sprintf("Olá %s,este é seu código para efetuarmos seu login", judgeModel.GetName())
	html := judge_util_service.EmailCode(title, msg, code)
	judge_util_service.SendEmailCode(twoAuth, title, code, journey, html)
	return nil
}
