package comment_response

type GetComment struct {
	CommentId	string `json:"comment_id"`
	ParentId	any 	 `json:"parent_id"`
	Answer	  any 	 `json:"answer_id"`
	Comment 	string `json:"comment"`
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	User      string `json:"user"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at"`
}