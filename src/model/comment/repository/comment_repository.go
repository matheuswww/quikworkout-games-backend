package comment_repository

import (
	"database/sql"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	comment_domain "github.com/matheuswww/quikworkout-games-backend/src/model/comment"
)

type commentRepository struct {
	mysql *sql.DB
}

type CommentRepository interface {
	CreateComment(comment comment_domain.CommentDomainInterface) *rest_err.RestErr
}

func NewCommentRepository(mysql *sql.DB) CommentRepository {
	return &commentRepository{
		mysql,
	}
}
