package comment_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	comment_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/comment/response"
)

func (cs *commentService) GetComment(video_id string, cursor, commentId string) ([]comment_response.GetComment, *rest_err.RestErr) {
	return cs.commentRepository.GetComment(video_id, cursor, commentId)
}