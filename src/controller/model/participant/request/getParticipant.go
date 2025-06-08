package participant_request

type GetParticipant struct {
	VideoId         string `form:"videoId"`
	EditionId       string `form:"editionId"`
	NotVideoId			string `form:"notVideoId"`
	CursorCreatedAt string `form:"cursorCreatedAt"`
	CursorUserTime  string `form:"cursorUserTime"`
	BestTime        bool   `form:"bestTime"`
	WorstTime       bool   `form:"worstTime"`
	Width           int    `form:"width"`
	Autoplay        bool   `form:"autoplay"`
	Muted           bool   `form:"muted"`
	Background      bool   `form:"background"`
}