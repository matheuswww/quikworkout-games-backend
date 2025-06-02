package edition_request

type GetEdition struct {
	Number int 		`form:"number"`
	Limit  int    `form:"limit"`
	Cursor string `form:"cursor"`
}