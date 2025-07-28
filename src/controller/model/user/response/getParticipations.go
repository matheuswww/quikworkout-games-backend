package user_response

type GetParticipations struct {
	Participations []Participantion `json:"participations"`
	User           User             `json:"user"`
	More           bool             `json:"more"`
}

type Participantion struct {
	Video        any    `json:"video"`
	Title        any    `json:"title"`
	ThumbnailUrl any    `json:"thumbnail_url"`
	VideoId      string `json:"video_id"`
	Sent         bool   `json:"sent"`
	Edition      int    `json:"edition"`
	EditionId    string `json:"edition_id"`
	Gain         any    `json:"gain"`
	Placing      any    `json:"placing"`
	FinalTime    any    `json:"final_time"`
	UserTime     string `json:"userTime"`
	Desqualified any    `json:"desqualified"`
	Category     string `json:"category"`
	Noreps       any    `json:"noreps"`
	Sex          string `json:"sex"`
	Checked      bool   `json:"checked"`
	CreatedAt    string `json:"createdAt"`
}

type User struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	User   string `json:"user"`
	Photo  string `json:"photo"`
}
