package user_response

type GetParticipations struct {
	Participations []Participantion `json:"participations"`
	User      		User     				  `json:"user"`
}

type Participantion struct {
	Video     		string   `json:"video"`
	Edition   		int      `json:"edition"`
	Placing   		any			 `json:"placing"`
	UserTime  		any   	 `json:"user_time"`
	Desqualified 	any   	 `json:"desqualified"`
	Checked       bool 		 `json:"checked"`
	CreatedAt 		string   `json:"createdAt"`
}

type User struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	User     string `json:"user"`
}

