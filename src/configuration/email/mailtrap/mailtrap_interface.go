package mailtrap

func NewMailTrap() Mailtrap {
	return &mailtrap{}
}

type mailtrap struct{}

type Mailtrap interface {
	NewMailTrapConnection(to, subject string, htmlContent []byte) error
}
