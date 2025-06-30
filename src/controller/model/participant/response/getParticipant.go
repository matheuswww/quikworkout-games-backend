package participant_response

type GetParticipant struct {
	Particiapants []Participant `json:"participants"`
	ClosingDate   string        `json:"closing_date"`
}

type Participant struct {
	Video        string `json:"video"`
	VideoId      string `json:"video_id"`
	Edition_id   string `json:"edition_id"`
	Category     string `json:"category"`
	Challenge    string `json:"challenge"`
	Title        any    `json:"title"`
	ThumbnailUrl any    `json:"thumbnail_url"`
	UserTime     any    `json:"user_time"`
	Placing      any    `json:"placing"`
	User         User   `json:"user"`
	CreatedAt    string `json:"createdAt"`
}

type User struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	User   string `json:"user"`
	Photo  string `json:"photo"`
}
