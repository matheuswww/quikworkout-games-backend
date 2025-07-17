package admin_request

type CreateEdition struct {
	StartDate    string      `json:"start_date" binding:"required" validate:"start_date"`
	ClosingDate  string      `json:"closing_date" binding:"required" validate:"closing_date"`
	ClothingName string      `json:"clothing_name" binding:"required,max=15"`
	Rules        string      `json:"rules" binding:"required"`
	Challenge    []Challenge `json:"challenge" binding:"required" validate:"dive"`
	Tops         []Top       `json:"tops" binding:"required" validate:"dive"`
}

type Top struct {
	Top      int    `json:"top" validate:"required"`
	Gain     int    `json:"gain"`
	Category string `json:"category" validate:"required,category"`
}

type Challenge struct {
	Challenge string `json:"challenge" validate:"required"`
	Category  string `json:"category" validate:"required,category"`
	Sex       string `json:"sex" validate:"required,sex"`
}