package email

func NewEmail() Email {
	return &email{}
}

type email struct{}

type Email interface {
	NewEmailConnection(to, subject string, htmlContent []byte) error
}
