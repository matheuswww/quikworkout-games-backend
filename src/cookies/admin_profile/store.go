package admin_profile_cookie

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
	encryptKey := os.Getenv("ADMIN_COOKIE_ENCRYPT")
	authKey := os.Getenv("ADMIN_COOKIE_AUTH")
	domain = os.Getenv("DOMAIN")
	envMode := os.Getenv("ENV_MODE")
	if domain == "" || envMode == "" || encryptKey == "" || authKey == "" {
		return nil, errors.New("error loading env")
	}
	secure = true
	if envMode == "DEV" {
		secure = false
	}
	store := cookie.NewStore([]byte(authKey), []byte(encryptKey))
	store.Options(sessions.Options{
		Path:     "/manager-quikworkout",
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
