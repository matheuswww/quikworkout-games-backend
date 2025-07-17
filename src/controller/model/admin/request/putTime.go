package admin_request

type PutTimeRequest struct {
	VideoId   string `json:"video_id" binding:"required" validate:"time"`
	EditionId string `json:"edition_id" binding:"required"`
	Category  string `json:"category" validate:"required,category"`
	Sex       string `json:"sex" validate:"required,sex"`
	Time      string `json:"time"`
}
