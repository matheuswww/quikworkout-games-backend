package admin_request

type CreateEdition struct {
	StartDate    string      `json:"start_date" binding:"required" validate:"start_date"`
	ClosingDate  string      `json:"closing_date" binding:"required" validate:"closing_date"`
	ClothingName string      `json:"clothing_name" binding:"required,max=15"`
	Rules        string      `json:"rules" binding:"required"`
	Challenge    []Challenge `json:"challenge" validate:"dive"`
	Tops         []Top       `json:"tops" binding:"required" validate:"dive"`
}

type Top struct {
	Top      int    `json:"top" binding:"required"`
	Gain     int    `json:"gain" binding:"required"`
	Category string `json:"category" binding:"required" validate:"category"`
}

type Challenge struct {
	Challenge string `json:"challenge" binding:"required"`
	Category  string `json:"category" binding:"required" validate:"category"`
}
