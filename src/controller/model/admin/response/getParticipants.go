package admin_response

type GetParticipants struct {
	Participants []Participant `json:"participants"`
}

type Participant struct {
	Video        any    `json:"video"`
	Title        any    `json:"title"`
	ThumbnailUrl any    `json:"thumbnail_url"`
	VideoId      string `json:"video_id"`
	Sent         bool   `json:"sent"`
	Edition      int    `json:"edition"`
	EditionId    string `json:"edition_id"`
	Placing      any    `json:"placing"`
	Gain 				 any    `json:"gain"`
	UserTime     any    `json:"user_time"`
	Desqualified any    `json:"desqualified"`
	Checked      bool   `json:"checked"`
	CreatedAt    string `json:"createdAt"`
	User         User   `json:"user"`
}

type User struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	User   string `json:"user"`
}
