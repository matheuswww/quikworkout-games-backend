package admin_request

type GetParticipants struct {
	EditionID       string `json:"edition_id"`
	CursorCreatedAt string `json:"cursor_createdAt"`
	CursorUserTime  string `json:"cursor_userTime"`
}