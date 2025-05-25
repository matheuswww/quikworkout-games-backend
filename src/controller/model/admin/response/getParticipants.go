package admin_response

type GetParticipants struct {
	Participants []Participant 	`json:"participations"`
}

type Participant struct {
	Video     		string   `json:"video"`
	Edition   		int      `json:"edition"`
	Placing   		any			 `json:"placing"`
	UserTime  		any   	 `json:"user_time"`
	Desqualified 	any   	 `json:"desqualified"`
	Checked       bool 		 `json:"checked"`
	CreatedAt 		string   `json:"createdAt"`
	User      		User     `json:"user"`
}

type User struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	User     string `json:"user"`
}

