package admin_request

type MakePlacing struct {
	EditionId string `json:"edition_id" binding:"required"`
	Category  string `json:"category" binding:"required" validate:"category"`
}
