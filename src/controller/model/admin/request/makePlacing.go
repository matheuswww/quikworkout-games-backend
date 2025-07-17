package admin_request

type MakePlacing struct {
	EditionId string `json:"edition_id" binding:"required"`
	Category  string `json:"category" validate:"required,category"`
	Sex       string `json:"sex" validate:"required,sex"`
}
