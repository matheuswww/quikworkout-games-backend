package admin_request

type DesqualifyVideo struct {
	VideoID      string `json:"video_id" binding:"required"`
	EditionId    string `json:"edition_id" binding:"required"`
	Category     string `json:"category" validate:"required,category"`
	Sex          string `json:"sex" validate:"required,sex"`
	Desqualified string `json:"desqualified"`
}
