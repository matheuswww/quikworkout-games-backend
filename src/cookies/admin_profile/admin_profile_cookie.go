package admin_profile_cookie

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"go.uber.org/zap"
)

var (
	SessionAdminProfile = "adminProfile"
)

type adminProfile struct {
	Id    string
	Email string
	Name  string
}

func SendAdminProfileCookie(c *gin.Context, id, name, email string) error {
	session := sessions.DefaultMany(c, SessionAdminProfile)
	adminProfile := adminProfile{
		Id:    id,
		Email: email,
		Name:  name,
	}
	session.Set("id", adminProfile.Id)
	session.Set("email", adminProfile.Email)
	session.Set("name", adminProfile.Name)
	SetOptions(session, "/manager-quikworkout", 3600*24*30)
	err := session.Save()
	if err != nil {
		return err
	}
	return nil
}

func GetAdminProfileValues(c *gin.Context) (cookie adminProfile, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Error trying get cookie", r.(error), zap.String("journey", "GetAdminProfileValues"))
			err = errors.New("invalid adminProfile cookie")
		}
	}()
	session := sessions.DefaultMany(c, SessionAdminProfile)
	id := session.Get("id")
	email := session.Get("email")
	name := session.Get("name")
	if id != nil && email != nil && name != nil {
		cookie = adminProfile{
			Id:    id.(string),
			Email: email.(string),
			Name:  name.(string),
		}
	} else {
		err = errors.New("invalid adminProfile cookie")
	}
	return
}
