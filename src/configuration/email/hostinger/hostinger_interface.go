package hostinger

func NewMailHostinger() Hostinger {
	return &hostinger{}
}

type hostinger struct{}

type Hostinger interface {
	NewMailHostingerConnection(to, subject string, htmlContent []byte) error
}
