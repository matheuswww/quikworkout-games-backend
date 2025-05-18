package hostinger

import (
	"errors"
	"net/smtp"
	"os"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
)

var (
	server = "smtp.hostinger.com"
	from   = "contact@quikworkout.com.br"
	port   = "587"
	user   = "contact@quikworkout.com.br"
)

func (m *hostinger) NewMailHostingerConnection(to, subject string, htmlContent []byte) error {
	password := os.Getenv("HOSTINGER_EMAIL_PASSWORD")
	if password == "" {
		logger.Error("Error trying get env", errors.New("hostinger email password not found"), zap.String("journey", "Hostinger"))
		return errors.New("env nout found")
	}
	emailHeaders := []byte(
		"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"From: " + from + "\r\n" +
			"Content-Type: text/html\r\n\r\n")

	emailContent := append(emailHeaders, htmlContent...)
	auth := smtp.PlainAuth("", user, password, server)
	err := smtp.SendMail(server+":"+port, auth, from, []string{to}, emailContent)
	if err != nil {
		return err
	}
	return nil
}
