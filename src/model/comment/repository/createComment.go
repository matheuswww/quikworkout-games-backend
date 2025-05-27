package comment_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	comment_domain "github.com/matheuswww/quikworkout-games-backend/src/model/comment"
	"go.uber.org/zap"
)

func (cr *commentRepository) CreateComment(comment comment_domain.CommentDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if comment.GetParentId() != "" {
		var video_id string
		query := "SELECT video_id FROM comment WHERE comment_id = ?"
		err := cr.mysql.QueryRowContext(ctx, query, comment.GetParentId()).Scan(&video_id)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error("Error comment not found", errors.New("comment not found"), zap.String("journey", "CreateComment Repository"))
				return rest_err.NewNotFoundError("parent comment not found")
			}
			logger.Error("Error trying QueryRowContext", err, zap.String("journey", "CreateComment Repository"))
			return rest_err.NewInternalServerError("server error")
		}
		comment.SetVideoId(video_id)
	} else {
		var count int
		query := "SELECT COUNT(*) FROM participant WHERE video_id = ?"
		err := cr.mysql.QueryRowContext(ctx, query, comment.GetVideoId()).Scan(&count)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error("Error video not found", errors.New("video not found"), zap.String("journey", "CreateComment Repository"))
				return rest_err.NewNotFoundError("parent comment not found")
			}
			logger.Error("Error trying QueryRowContext", err, zap.String("journey", "CreateComment Repository"))
			return rest_err.NewInternalServerError("server error")
		}
		if count == 0 {
			logger.Error("Error video not found", errors.New("video not found"), zap.String("journey", "CreateComment Repository"))
			return rest_err.NewNotFoundError("video not found")
		}
	}
	
	commentId := uuid.NewString()
	comment.SetCommentId(commentId)
	var parentId any = comment.GetParentId()
	if comment.GetParentId() == "" {
		parentId = nil
	}
	query := "INSERT INTO comment (comment_id, video_id, parent_id, user_id, video_comment) VALUES (?, ?, ?, ?, ?)"
	_,err := cr.mysql.ExecContext(ctx, query, comment.GetCommentId(), comment.GetVideoId(),  parentId, comment.GetUserId(), comment.GetVideoComment())
	if err != nil {
		logger.Error("Error trying ExecContext", err, zap.String("journey", "CreateComment Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	return nil
}