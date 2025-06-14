package user_response

type GetAccount struct {
	Name 		 string `json:"name"`
	User 		 string `json:"user"`
	Category string `json:"category"`
	Earnings int 		`json:"earnings"`
	Photo    string `json:"photo"`
}