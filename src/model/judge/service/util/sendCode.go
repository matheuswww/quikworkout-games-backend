package judge_util_service

import (
	"fmt"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/email"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
)

func SendEmailCode(to, subject, code, journey string, html []byte) {
	email := email.NewEmail()
	go func() {
		err := email.NewEmailConnection(to, subject, html)
		if err != nil {
			logger.Error("Error trying send email", err, zap.String("journey", journey))
		}
		logger.Info(fmt.Sprintf("Email sended with success, subject: %s, to: %s", subject, to), zap.String("journey", journey))
	}()
}
