package recaptcha

import "github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"

func NewRecaptcha() Recaptcha {
	return &recaptcha{}
}

type recaptcha struct{}

type Recaptcha interface {
	ValidateRecaptcha(token string) *rest_err.RestErr
}
