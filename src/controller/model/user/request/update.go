package user_request

import "mime/multipart"

type Update struct {
	User     string                `form:"user" binding:"omitempty,min=1,max=30" validate:"user"`
	Name     string                `form:"name" binding:"omitempty,min=2,max=20"`
	Category string                `form:"category" binding:"omitempty,max=10"`
	Token    string                `form:"token" binding:"max=10000"`
	Image    *multipart.FileHeader `form:"imagem"`
}
