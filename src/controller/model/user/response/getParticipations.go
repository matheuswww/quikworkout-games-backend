package user_response

type GetParticipations struct {
	Participations []Participantion `json:"participations"`
	User      		User     				  `json:"user"`
}

type Participantion struct {
	Video     		any   `json:"video"`
	Title         any   `json:"title"`
	ThumbnailUrl  any 	`json:"thumbnail_url"`
	Edition   		int   `json:"edition"`
	Gain   				any   `json:"gain"`
	Placing   		any		`json:"placing"`
	UserTime  		any   `json:"user_time"`
	Desqualified 	any   `json:"desqualified"`
	Checked       bool 	`json:"checked"`
	CreatedAt 		string `json:"createdAt"`
}

type User struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	User     string `json:"user"`
}

