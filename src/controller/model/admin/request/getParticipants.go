package admin_request

type GetParticipants struct {
	EditionID       string `form:"edition_id"`
	CursorCreatedAt string `form:"cursor_createdAt"`
	CursorUserTime  string `form:"cursor_userTime"`
}