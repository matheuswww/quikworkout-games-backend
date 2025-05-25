package admin_request

type CheckVideo struct {
	VideoID string `json:"video_id" binding:"required"`
}