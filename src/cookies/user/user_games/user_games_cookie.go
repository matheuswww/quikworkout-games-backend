package user_games_cookie

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
	SessionUserGames = "userGames"
)

type UserGamesCookie struct {
	Id        string
	SessionId string
	Games     bool
}

func SendUserGamesCookie(c *gin.Context, id, sessionId string, games bool) error {
	session := sessions.DefaultMany(c, SessionUserGames)
	sessionCookie := UserGamesCookie{
		Id:        id,
		SessionId: sessionId,
		Games: games,
	}
	session.Set("id", sessionCookie.Id)
	session.Set("sessionId", sessionCookie.SessionId)
	session.Set("games", sessionCookie.Games)
	user_cookie.SetOptions(session, "/", 3600*24*30)
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}

func GetUserGamesCookieValues(c *gin.Context) (cookie UserGamesCookie, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Error trying get cookie", r.(error), zap.String("journey", "GetUserGamesCookieValues"))
			err = errors.New("invalid UserGamesCookie")
		}
	}()
	session := sessions.DefaultMany(c, SessionUserGames)
	id := session.Get("id")
	sessionId := session.Get("sessionId")
	games := session.Get("games")
	if id != nil && sessionId != nil && games != nil {
		cookie = UserGamesCookie{
			Id:        id.(string),
			SessionId: sessionId.(string),
			Games: games.(bool),
		}
	} else {
		err = errors.New("invalid UserGamesCookie")
	}
	return
}

func Clear(c *gin.Context) {
	session := sessions.DefaultMany(c, SessionUserGames)
	session.Clear()
	domain := os.Getenv("DOMAIN")
	env := os.Getenv("ENV_MODE")
	var secure = true
	if env == "DEV" {
		secure = false
	}
	c.SetCookie(SessionUserGames, "", -1, "/", domain, secure, true)
}
