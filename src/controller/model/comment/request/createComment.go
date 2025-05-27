package comment_request

type CreateComment struct {
	VideoComment string `json:"video_comment" binding:"required,max=100"`
	VideoId      string `json:"video_id"`
	ParentId     string `json:"parent_id"`
}