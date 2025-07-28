package participant_request

type GetParticipant struct {
	VideoId         string `form:"videoId"`
	EditionId       string `form:"editionId"`
	NotVideoId      string `form:"notVideoId"`
	CursorPlacing   int    `form:"cursorPlacing"`
	Category        string `form:"category" validate:"category"`
	Sex             string `form:"sex" validate:"sex"`
	BestTime        bool   `form:"bestTime"`
	Width           int    `form:"width"`
	Autoplay        bool   `form:"autoplay"`
	Muted           bool   `form:"muted"`
	Background      bool   `form:"background"`
}
