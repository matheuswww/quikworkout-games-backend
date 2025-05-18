package mailtrap

import (
	"errors"
	"net/smtp"
	"os"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
)

var (
	server = "sandbox.smtp.mailtrap.io"
	from   = "test@test.com"
	port   = "587"
	user   = "f46a527bf58ca4"
)

func (m *mailtrap) NewMailTrapConnection(to, subject string, htmlContent []byte) error {
	password := os.Getenv("MAILTRAP_PASSWORD")
	if password == "" {
		logger.Error("Error trying get env", errors.New("mailtrap password not found"), zap.String("journey", "MailTrap"))
		return errors.New("env nout found")
	}
	emailHeaders := []byte(
		"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"From: " + from + "\r\n" +
			"Content-Type: text/html\r\n\r\n")

	emailContent := append(emailHeaders, htmlContent...)
	auth := smtp.CRAMMD5Auth(user, password)
	err := smtp.SendMail(server+":"+port, auth, from, []string{to}, emailContent)
	if err != nil {
		return err
	}
	return nil
}
