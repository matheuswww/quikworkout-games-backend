package admin_request

type GrantTicket struct {
	User string `json:"user" binding:"required,min=1,max=30"`
}