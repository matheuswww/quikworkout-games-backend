package user_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/request"
	user_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/response"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	user_repository "github.com/matheuswww/quikworkout-games-backend/src/model/user/repository"
)


func NewUserService(userRepository user_repository.UserRepository) UserService {
	return &userService{
		userRepository,
	}
}

type userService struct {
	userRepository user_repository.UserRepository
}

type UserService interface {
	CreateAccount(userDomain user_domain.UserDomainInterface, saveImg func() error, id, token string) *rest_err.RestErr
	EnterAccount(userDomain user_domain.UserDomainInterface) *rest_err.RestErr
	GetAccount(userDomain user_domain.UserDomainInterface, sessionId string) *rest_err.RestErr
	GetParticipations(user_domain user_domain.UserDomainInterface, getParticipartRequest *user_request.GetParticipations) (*user_response.GetParticipations, *rest_err.RestErr)
}
