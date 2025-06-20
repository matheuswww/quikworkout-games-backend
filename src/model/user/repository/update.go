package user_repository

import (
	"context"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/request"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	"go.uber.org/zap"
)

func (ur *userRepository) Update(userDomain user_domain.UserDomainInterface, updateRequest *user_request.Update) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if updateRequest.User == "" && updateRequest.Category == "" && updateRequest.Name == "" {
		return rest_err.NewBadRequestError("invalid params")
	} 
	if updateRequest.User != "" {
		var count int
		query := "SELECT COUNT(*) FROM user_games WHERE user = ?"
		err := ur.mysql.QueryRowContext(ctx, query, updateRequest.User).Scan(&count)
		if err != nil {
			logger.Error("Error trying Update", err, zap.String("journey", "Update Repository"))
			return rest_err.NewInternalServerError("server error")
		}
		if count > 0 {
			return rest_err.NewBadRequestError("user already exists")
		}
	}
	var args []any
	query :=  "UPDATE user_games SET "
	if updateRequest.User != "" {
		query += "user = ?,"
		args = append(args, updateRequest.User)
	}
	if updateRequest.Category != "" {
		query += "category = ?,"
		args = append(args, updateRequest.Category)
	}
	if updateRequest.Name != "" {
		query += "name = ?,"
		args = append(args, updateRequest.Name)
	}
	query = query[:len(query)-1]
	query += " WHERE user_id = ? "
	args = append(args, userDomain.GetId())
	_,err := ur.mysql.ExecContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying ExecContext", err, zap.String("journey", "Update Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}