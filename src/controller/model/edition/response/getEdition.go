package edition_response

type Top struct {
	Gain     int    `json:"gain"`
	Top      int    `json:"top"`
	Category string `json:"category"`
}

type Challenge struct {
	Challenge string `json:"challenge"`
	Category  string `json:"category"`
}

type Edition struct {
	Id          string      `json:"id"`
	StartDate   string      `json:"start_date"`
	ClosingDate string      `json:"closing_date"`
	Rules       string      `json:"rules"`
	Number      int         `json:"number"`
	CreatedAt   string      `json:"created_at"`
	Challenge   []Challenge `json:"challenges"`
	Tops        []Top       `json:"tops"`
}
