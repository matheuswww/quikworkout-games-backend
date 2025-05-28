package edition_response

type Top struct {
	Gain int `json:"gain"`
	Top  int `json:"top"`
}

type Edition struct {
	Id					string `json:"id"`
	StartDate   string `json:"start_date"`
	ClosingDate string `json:"closing_date"`
	Rules       string `json:"rules"`
	Challenge   string `json:"challenge"`
	Number 			int    `json:"number"`
	CreatedAt   string `json:"created_at"`
	Tops        []Top  `json:"tops"`
}