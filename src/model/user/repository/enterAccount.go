package user_repository

import (
	"context"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	"go.uber.org/zap"
)

func (ur *userRepository) EnterAccount(userDomain user_domain.UserDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	query := "UPDATE user_games SET session_id = ? WHERE user_id = ?"
	_,err := ur.mysql.ExecContext(ctx, query, userDomain.GetSessionId(), userDomain.GetId())
	if err != nil {
		logger.Error("Error trying ExecContext", err, zap.String("journey", "EnterAccount Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}