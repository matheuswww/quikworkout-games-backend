package admin_profile_cookie

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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

func GetAdminProfileValues(c *gin.Context) (adminProfile, error) {
	session := sessions.DefaultMany(c, SessionAdminProfile)
	id := session.Get("id")
	email := session.Get("email")
	name := session.Get("name")
	if id != nil && email != nil && name != nil {
		return adminProfile{
			Id:    id.(string),
			Email: email.(string),
			Name:  name.(string),
		}, nil
	}
	return adminProfile{}, errors.New("invalid cookie admin profile")
}
