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

func (ur *userRepository) CreateAccount(userDomain user_domain.UserDomainInterface, sessionIdFromCookie string, saveImg func() *rest_err.RestErr) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var name, session_id string
	query := "SELECT name, session_id FROM user WHERE user_id = ?"
	err := ur.mysql.QueryRowContext(ctx, query, userDomain.GetId()).Scan(&name, &session_id)
	if err != nil {
		logger.Error("Error trying get user", err, zap.String("journey", "CreateAccount Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if sessionIdFromCookie != session_id {
		logger.Error("Error invalid session", errors.New("invalid session"), zap.String("journey", "ChangePassword Repository"))
		return rest_err.NewUnauthorizedError("cookie inválido")
	}

	query = "SELECT COUNT(*) FROM user_games WHERE user_id = ?"
	var count int
	err = ur.mysql.QueryRowContext(ctx, query, userDomain.GetId()).Scan(&count)
	if err != nil {
		logger.Error("Error trying get user_games", err, zap.String("journey", "CreateAccount Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if count != 0 {
		logger.Error("Error account already created", errors.New("account already created"), zap.String("journey", "CreateAccount Repository"))
		return rest_err.NewBadRequestError("conta já criada")
	}
	userDomain.SetName(name)

	query = "SELECT COUNT(*) FROM user_games WHERE user = ?"
	err = ur.mysql.QueryRowContext(ctx, query, userDomain.GetUser()).Scan(&count)
	if err != nil {
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "CreateAccount Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if count != 0 {
		logger.Error("Error user already exists", errors.New("user already exists"), zap.String("journey", "CreateAccount Repository"))
		return rest_err.NewBadRequestError("usuário já cadastrado")
	}

	restErr := saveImg()
	if restErr != nil {
		return restErr
	}

	query = "INSERT INTO user_games (user_id, name, user, category, earnings, session_id) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = ur.mysql.ExecContext(ctx, query, userDomain.GetId(), userDomain.GetName(), userDomain.GetUser(), userDomain.GetCategory(), userDomain.GetEarnings(), userDomain.GetSessionId())
	if err != nil {
		logger.Error("Error trying insert user", err, zap.String("journey", "CreateAccount Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}