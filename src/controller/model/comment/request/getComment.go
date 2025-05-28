package comment_request

type GetComment struct {
	VideoId   string `form:"video_id" binding:"required,max=36"`
	CommentId string `form:"comment_id"`
	Cursor		string `form:"cursor"`
}