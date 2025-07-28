package admin_response

type GetParticipants struct {
	Participants []Participant `json:"participants"`
	ClosingDate  string        `json:"closing_date"`
	More         bool          `json:"more"`
}

type Participant struct {
	Video        any    `json:"video"`
	Title        any    `json:"title"`
	ThumbnailUrl any    `json:"thumbnail_url"`
	VideoId      string `json:"video_id"`
	Category     string `json:"category"`
	Noreps       any    `json:"noreps"`
	Sex          string `json:"sex"`
	Sent         bool   `json:"sent"`
	Edition      int    `json:"edition"`
	EditionId    string `json:"edition_id"`
	Challenge    string `json:"challenge"`
	UserTime     string `json:"user_time"`
	Placing      any    `json:"placing"`
	Gain         any    `json:"gain"`
	FinalTime    any    `json:"final_time"`
	Desqualified any    `json:"desqualified"`
	Checked      bool   `json:"checked"`
	CreatedAt    string `json:"createdAt"`
	User         User   `json:"user"`
}

type User struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	User   string `json:"user"`
	Email  string `json:"email"`
	Photo  string `json:"photo"`
}
