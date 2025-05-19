package user_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/recaptcha"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
)


func (us *userService) CreateAccount(userDomain user_domain.UserDomainInterface, id, token string) *rest_err.RestErr {
	logger.Info("Init CreateAccount service")
	recaptchaErr := recaptcha.NewRecaptcha().ValidateRecaptcha(token)
	if recaptchaErr != nil {
		return recaptchaErr
	}
	return us.userRepository.CreateAccount(userDomain, id)
}