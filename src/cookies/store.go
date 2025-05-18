package user_cookie

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var (
	domain   string
	secure   bool
)

func Store() (cookie.Store, error) {
	encryptKey := os.Getenv("COOKIE_ENCRYPT")
	authKey := os.Getenv("COOKIE_AUTH")
	domain = os.Getenv("DOMAIN")
	env := os.Getenv("ENV_MODE")
	if domain == "" || env == "" || encryptKey == "" || authKey == "" {
		return nil, errors.New("error loading env")
	}
	secure = true
	if env == "DEV" {
		secure = false
	}
	store := cookie.NewStore([]byte(authKey), []byte(encryptKey))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   (3600 * 24) * 30,
		HttpOnly: true,
		Secure:   secure,
		Domain:   domain,
		SameSite: http.SameSiteLaxMode,
	})
	return store, nil
}

func SetOptions(session sessions.Session, path string, exp int) {
	session.Options(sessions.Options{
		Path:     path,
		MaxAge:   exp,
		HttpOnly: true,
		Secure:   secure,
		Domain:   domain,
		SameSite: http.SameSiteLaxMode,
	})
}
