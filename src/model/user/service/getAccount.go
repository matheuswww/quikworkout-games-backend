package user_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
)

func (us *userService) GetAccount(userDomain user_domain.UserDomainInterface, sessionId string) *rest_err.RestErr {
	return us.userRepository.GetAccount(userDomain, sessionId)
}