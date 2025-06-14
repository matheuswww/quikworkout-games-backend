package admin_request

type PutTimeRequest struct {
	VideoId   string `json:"video_id" binding:"required" validate:"time"`
	EditionId string `json:"edition_id" binding:"required"`
	Time      string `json:"time"`
}
