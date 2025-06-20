package user_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/request"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
)

func (us *userService) Update(userDomain user_domain.UserDomainInterface, updateRequest *user_request.Update) *rest_err.RestErr {
	return us.userRepository.Update(userDomain, updateRequest)
}