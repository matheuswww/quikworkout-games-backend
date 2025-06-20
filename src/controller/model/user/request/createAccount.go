package user_request

import "mime/multipart"

type CreateAccount struct {
	User     string                `form:"user" binding:"required,min=1,max=30" validate:"user"`
	Category string                `form:"category" binding:"required,max=10"`
	Image    *multipart.FileHeader `form:"imagem" binding:"required"`
	Token    string                `form:"token" binding:"required,max=10000"`
}
