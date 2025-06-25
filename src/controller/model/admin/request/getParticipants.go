package admin_request

type GetParticipants struct {
	VideoId         string `form:"videoId"`
	EditionId       string `form:"editionId"`
	CursorCreatedAt string `form:"cursorCreatedAt"`
	CursorUserTime  string `form:"cursorUserTime"`
	Category        string `form:"category" validate:"category"`
	BestTime        bool   `form:"bestTime"`
	Width           int    `form:"width"`
	Autoplay        bool   `form:"autoplay"`
	Muted           bool   `form:"muted"`
	Background      bool   `form:"background"`
}
