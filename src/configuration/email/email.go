package email

import (
	"context"
	"os"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/email/hostinger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/email/mailtrap"
)

func (m *email) NewEmailConnection(to, subject string, htmlContent []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	errCh := make(chan error, 1)
	go func() {
		var err error
		var mode string = os.Getenv("ENV_MODE")
		if mode == "DEV" {
			mailtrap := mailtrap.NewMailTrap()
			err = mailtrap.NewMailTrapConnection(to, subject, htmlContent)
		} else if mode == "PROD" {
			hostinger := hostinger.NewMailHostinger()
			err = hostinger.NewMailHostingerConnection(to, subject, htmlContent)
		}
		errCh <- err
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errCh:
		return err
	}
}
