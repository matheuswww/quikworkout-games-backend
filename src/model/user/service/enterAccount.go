package user_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
)

func (us *userService) EnterAccount(userDomain user_domain.UserDomainInterface) *rest_err.RestErr {
	return us.userRepository.EnterAccount(userDomain)
}