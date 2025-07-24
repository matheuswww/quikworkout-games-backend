package admin_request

import "mime/multipart"

type CreateEdition struct {
	StartDate    string                `form:"start_date" binding:"required" validate:"start_date"`
	ClosingDate  string                `form:"closing_date" binding:"required" validate:"closing_date"`
	ClothingName string                `form:"clothing_name" binding:"required,max=15"`
	Rules        *multipart.FileHeader `form:"rules" binding:"required"`
	Challenge    string           	   `form:"challenge" binding:"required"`
	Tops         string                `form:"tops" binding:"required"`
}

type Top struct {
	Top      int
	Gain     int
	Category string
}

type Challenge struct {
	Challenge string
	Category  string
	Sex       string
}
