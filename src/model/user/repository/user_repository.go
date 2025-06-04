package user_repository

import (
	"database/sql"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
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
	GetParticipations(user_domain user_domain.UserDomainInterface, limit int, cursor string) (*user_response.GetParticipations, *sql.DB, *rest_err.RestErr)
}
