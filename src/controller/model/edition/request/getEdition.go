package edition_request

type GetEdition struct {
	Number int 		`form:"number"`
	Cursor string `form:"cursor"`
}