package comment_service

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	comment_domain "github.com/matheuswww/quikworkout-games-backend/src/model/comment"
	comment_repository "github.com/matheuswww/quikworkout-games-backend/src/model/comment/repository"
)

type commentService struct {
	commentRepository comment_repository.CommentRepository
}

type CommentService interface {
	CreateComment(comment comment_domain.CommentDomainInterface) *rest_err.RestErr
}

func NewCommentService(commentRepository comment_repository.CommentRepository) CommentService {
	return &commentService{
		commentRepository,
	}
}
