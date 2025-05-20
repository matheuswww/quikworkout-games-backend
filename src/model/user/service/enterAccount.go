package user_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	user_service_util "github.com/matheuswww/quikworkout-games-backend/src/model/user/service/util"
	"go.uber.org/zap"
)

func (us *userService) EnterAccount(userDomain user_domain.UserDomainInterface) *rest_err.RestErr {
	sessionId, sessionErr := user_service_util.GenerateNewSessionId()
	if sessionErr != nil {
		logger.Error("Error trying generate new session id", sessionErr, zap.String("journey", "CreateAccount Service"))
		return rest_err.NewInternalServerError("server error")
	}
	userDomain.SetSessionId(sessionId)
	return us.userRepository.EnterAccount(userDomain)
}