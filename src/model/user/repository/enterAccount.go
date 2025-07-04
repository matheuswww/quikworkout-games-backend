package user_repository

import (
	"context"
	"database/sql"
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
	var session_id string
	query := "SELECT session_id FROM user_games WHERE user_id = ?"
	err := ur.mysql.QueryRowContext(ctx, query, userDomain.GetId()).Scan(&session_id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error User not found", errors.New("user not found"), zap.String("journey", "EnterAccount Repository"))
			return rest_err.NewBadRequestError("usuário não encontrado")
		}
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "EnterAccount Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	userDomain.SetSessionId(session_id)
	return nil
}