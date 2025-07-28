package judge_profile_cookie

import (
	"errors"

	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	user_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user"
	"go.uber.org/zap"
)

var (
	SessionJudge = "judgeProfile"
)

type JudgeCookie struct {
	Id        string
}

func SendJudgeCookie(c *gin.Context, id string) error {
	session := sessions.DefaultMany(c, SessionJudge)
	sessionCookie := JudgeCookie{
		Id:        id,
	}
	session.Set("id", sessionCookie.Id)
	user_cookie.SetOptions(session, "/judge", 3600*24*30)
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}

func GetJudgeCookieValues(c *gin.Context) (cookie JudgeCookie, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Error trying get cookie", r.(error), zap.String("journey", "GetJudgeCookieValues"))
			err = errors.New("invalid JudgeCookie")
		}
	}()
	session := sessions.DefaultMany(c, SessionJudge)
	id := session.Get("id")
	if id != nil {
		cookie = JudgeCookie{
			Id: id.(string),
		}
	} else {
		err = errors.New("invalid JudgeCookie")
	}
	return
}

func Clear(c *gin.Context) {
	session := sessions.DefaultMany(c, SessionJudge)
	session.Clear()
	domain := os.Getenv("DOMAIN")
	env := os.Getenv("ENV_MODE")
	var secure = true
	if env == "DEV" {
		secure = false
	}
	c.SetCookie(SessionJudge, "", -1, "/", domain, secure, true)
}
