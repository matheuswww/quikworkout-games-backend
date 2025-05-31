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

func (ur *userRepository) EnterAccount(userDomain user_domain.UserDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	query := "SELECT COUNT(*) FROM user_games WHERE user_id = ?"
	var count int
	err := ur.mysql.QueryRowContext(ctx, query, userDomain.GetId()).Scan(&count)
	if err != nil {
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "EnterAccount Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if count == 0 {
		logger.Error("Error User not found", errors.New("user not found"), zap.String("journey", "EnterAccount Repository"))
		return rest_err.NewBadRequestError("usuário não encontrado")
	}
	query = "UPDATE user_games SET session_id = ? WHERE user_id = ?"
	_,err = ur.mysql.ExecContext(ctx, query, userDomain.GetSessionId(), userDomain.GetId())
	if err != nil {
		logger.Error("Error trying ExecContext", err, zap.String("journey", "EnterAccount Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}