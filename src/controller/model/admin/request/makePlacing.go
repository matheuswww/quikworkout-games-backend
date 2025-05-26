package admin_request

type MakePlacing struct {
	EditionId string `json:"edition_id" binding:"required"`
}