package user_games_cookie

import (
	"errors"

	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	user_cookie "github.com/matheuswww/quikworkout-games-backend/src/cookies/user"
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

func GetUserGamesCookieValues(c *gin.Context) (UserGamesCookie, error) {
	session := sessions.DefaultMany(c, SessionUserGames)
	id := session.Get("id")
	sessionId := session.Get("sessionId")
	games := session.Get("games")
	if id != nil && sessionId != nil && games != nil {
		return UserGamesCookie{
			Id:        id.(string),
			SessionId: sessionId.(string),
			Games: games.(bool),
		}, nil
	}
	return UserGamesCookie{}, errors.New("invalid UserGamesCookie")
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
