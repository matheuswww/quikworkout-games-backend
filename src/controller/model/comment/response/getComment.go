package comment_response

type CreateComment struct {
	CommentId	string `json:"comment_id"`
	Comment 	string `json:"comment"`
	ParentId	string `json:"parent_id"`
	User 		  string `json:"user"`
}