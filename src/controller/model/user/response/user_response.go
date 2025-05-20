package user_response

type GetAccount struct {
	Name 		 string `json:"name"`
	User 		 string `json:"user"`
	Dob 	 	 string `json:"dob"`
	Category string `json:"category"`
	Earnings int 		`json:"earnings"`
}