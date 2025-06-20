package user_repository

import (
	"database/sql"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/request"
	user_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/response"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
)

func NewUserRepository(mysql *sql.DB) UserRepository {
	return &userRepository{
		mysql,
	}
}

type userRepository struct {
	mysql *sql.DB
}

type UserRepository interface {
	CreateAccount(userDomain user_domain.UserDomainInterface, sessionIdFromCookie string) *rest_err.RestErr
	EnterAccount(userDomain user_domain.UserDomainInterface) *rest_err.RestErr
	GetAccount(userDomain user_domain.UserDomainInterface, sessionIdFromCookie string) *rest_err.RestErr
	GetParticipations(user_domain user_domain.UserDomainInterface, getParticipationsRequest *user_request.GetParticipations) (*user_response.GetParticipations, *sql.DB, *rest_err.RestErr)
	Update(userDomain user_domain.UserDomainInterface, updateRequest *user_request.Update) *rest_err.RestErr
}
