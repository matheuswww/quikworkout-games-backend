package user_repository

import (
	"context"
	"errors"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	"go.uber.org/zap"
)

func (ur *userRepository) GetAccount(userDomain user_domain.UserDomainInterface, sessionIdFromCookie string) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var name, user, category, session_id string
	var earnings int
	query := "SELECT name, user, category, earnings, session_id FROM user_games WHERE user_id = ?"
	err := ur.mysql.QueryRowContext(ctx, query, userDomain.GetId()).Scan(&name, &user, &category, &earnings, &session_id)
	if err != nil {
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetAccount Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if sessionIdFromCookie != session_id {
		logger.Error("Error invalid session", errors.New("invalid session"), zap.String("journey", "GetAccount Repository"))
		return rest_err.NewUnauthorizedError("cookie inválido")
	}
	userDomain.SetName(name)
	userDomain.SetUser(user)
	userDomain.SetCategory(category)
	userDomain.SetEarnings(earnings)
	return nil
}