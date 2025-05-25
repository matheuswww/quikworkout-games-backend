package user_request

type GetParticipations struct {
	Cursor string `form:"cursor"`
}