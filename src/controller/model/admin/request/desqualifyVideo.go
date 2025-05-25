package admin_request

type DesqualifyVideo struct {
	VideoID      string `json:"video_id" binding:"required"`
	Desqualified string `json:"desqualified"`
}