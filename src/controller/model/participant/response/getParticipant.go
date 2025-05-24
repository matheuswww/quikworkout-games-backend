package participant_response

type GetParticipant struct {
	Particiapants []Participant `json:"participants"`
}

type Participant struct {
	Video     string   `json:"video"`
	User      User     `json:"user"`
	CreatedAt string   `json:"createdAt"`
}

type User struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	User     string `json:"user"`
	UserTime any    `json:"userTime"`
}

