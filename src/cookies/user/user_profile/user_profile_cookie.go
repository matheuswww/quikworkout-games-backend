package user_proflie_cookie

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
	SessionUserProfile = "userProfile"
)

type userProfileCookie struct {
	Id        string
	SessionId string
}

func SendUserProfileCookie(c *gin.Context, id, sessionId string) error {
	session := sessions.DefaultMany(c, SessionUserProfile)
	sessionCookie := userProfileCookie{
		Id:        id,
		SessionId: sessionId,
	}
	session.Set("id", sessionCookie.Id)
	session.Set("sessionId", sessionCookie.SessionId)
	user_cookie.SetOptions(session, "/", 3600*24*30)
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}

func GetUserProfileCookieValues(c *gin.Context) (cookie userProfileCookie, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Error trying get cookie", r.(error), zap.String("journey", "GetUserProfileCookieValues"))
			err = errors.New("invalid userProfileCookie")
		}
	}()
	session := sessions.DefaultMany(c, SessionUserProfile)
	id := session.Get("id")
	sessionId := session.Get("sessionId")
	if id != nil && sessionId != nil {
		cookie = userProfileCookie{
			Id:        id.(string),
			SessionId: sessionId.(string),
		}
	} else {
		err =  errors.New("invalid userProfileCookie")
	}
	return
}

func Clear(c *gin.Context) {
	session := sessions.DefaultMany(c, SessionUserProfile)
	session.Clear()
	domain := os.Getenv("DOMAIN")
	env := os.Getenv("ENV_MODE")
	var secure = true
	if env == "DEV" {
		secure = false
	}
	c.SetCookie(SessionUserProfile, "", -1, "/", domain, secure, true)
}
