package user_request

type GetParticipations struct {
	Cursor string   `form:"cursor"`
	Width      int  `form:"width"`
	Autoplay   bool `form:"autoplay"`
	Muted      bool	`form:"muted"`
	Background bool	`form:"background"`
	Limit      int  `form:"limit"`
}