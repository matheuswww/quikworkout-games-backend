package user_request

type GetParticipations struct {
	Cursor string     `form:"cursor"`
	Width      int    `form:"width"`
	Autoplay   bool   `form:"autoplay"`
	Muted      bool	  `form:"muted"`
	Background bool	  `form:"background"`
	VideoId    string `form:"video_id"`
	EditionId  string `form:"edition_id"`
	Limit      int    `form:"limit"`
}