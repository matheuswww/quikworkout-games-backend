package user_repository

import (
	"database/sql"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
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
}
