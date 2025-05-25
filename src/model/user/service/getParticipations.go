package user_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/response"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
)

func (us *userService) GetParticipations(user_domain user_domain.UserDomainInterface, cursor string) (*user_response.GetParticipations, *rest_err.RestErr) {
	return us.userRepository.GetParticipations(user_domain, cursor)
}