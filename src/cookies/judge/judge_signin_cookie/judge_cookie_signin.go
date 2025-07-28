package judge_signin_cookie

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	user_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user"
	"go.uber.org/zap"
)

var (
	SessionSignin = "judgeSignin"
)

type judgeSignin struct {
	Id  string
}

func SendSigninCookie(c *gin.Context, id string) error {
	session := sessions.DefaultMany(c, SessionSignin)
	judgeSignin := judgeSignin{
		Id:    id,
	}
	session.Set("id", judgeSignin.Id)
	user_cookie.SetOptions(session, "/judge/auth", 3*60)
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}

func GetSigninValues(c *gin.Context) (cookie judgeSignin, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Error trying get cookie", r.(error), zap.String("journey", "GetSigninValues"))
			err = errors.New("invalid judgeSignin cookie")
		}
	}()
	session := sessions.DefaultMany(c, SessionSignin)
	id := session.Get("id")
	if id != nil {
		cookie = judgeSignin{
			Id:    id.(string),
		}
	} else {
		err = errors.New("invalid judgeSignin cookie")
	}
	return
}
