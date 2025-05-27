package comment_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	comment_domain "github.com/matheuswww/quikworkout-games-backend/src/model/comment"
)

func (cs *commentService) CreateComment(comment comment_domain.CommentDomainInterface) *rest_err.RestErr {
	return cs.commentRepository.CreateComment(comment)
}