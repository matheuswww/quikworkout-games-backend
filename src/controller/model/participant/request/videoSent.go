package participant_request

type VideoSent struct {
	VideoId string `json:"video_id" binding:"required"`
}