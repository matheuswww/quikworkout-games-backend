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

	if comment.GetAnswerId() != "" && comment.GetParentId() == "" {
		logger.Error("Error trying to create comment", errors.New("parent_id is empty"), zap.String("journey", "CreateComment Repository"))
		return rest_err.NewBadRequestError("parent_id is empty")
	}

	if comment.GetParentId() != "" {
		var video_id, comment_id string
		var parent_id sql.NullString 
		query := "SELECT video_id, comment_id, parent_id FROM comment WHERE comment_id = ?"
		err := cr.mysql.QueryRowContext(ctx, query, comment.GetParentId()).Scan(&video_id, &comment_id, &parent_id)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error("Error comment not found", errors.New("comment not found"), zap.String("journey", "CreateComment Repository"))
				return rest_err.NewNotFoundError("parent comment not found")
			}
			logger.Error("Error trying QueryRowContext", err, zap.String("journey", "CreateComment Repository"))
			return rest_err.NewInternalServerError("server error")
		}
		if parent_id.Valid {
			logger.Error("Error trying to create comment", errors.New("invalid parent_id"), zap.String("journey", "CreateComment Repository"))
			return rest_err.NewBadRequestError("invalid parent_id")
		} 
		if comment.GetAnswerId() != "" {
			var user_id string
			var parent_id sql.NullString
			query := "SELECT user_id, parent_id FROM comment WHERE comment_id = ?"
			err := cr.mysql.QueryRowContext(ctx, query, comment.GetAnswerId()).Scan(&user_id, &parent_id)
			if err != nil {
				if err == sql.ErrNoRows {
					logger.Error("Error comment not found", errors.New("comment not found"), zap.String("journey", "CreateComment Repository"))
					return rest_err.NewNotFoundError("answer comment not found")
				}
				logger.Error("Error trying QueryRowContext", err, zap.String("journey", "CreateComment Repository"))
				return rest_err.NewInternalServerError("server error")
			}
			if user_id == comment.GetUserId() {
				logger.Error("Error trying to create comment", errors.New("invalid answer_id"), zap.String("journey", "CreateComment Repository"))
				return rest_err.NewBadRequestError("invalid answer_id")
			}
			if parent_id.Valid && parent_id.String != comment.GetParentId() {
				logger.Error("Error trying to create comment", errors.New("invalid answer_id"), zap.String("journey", "CreateComment Repository"))
				return rest_err.NewBadRequestError("invalid answer_id")
			}
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
	var answerId any = comment.GetAnswerId()
	if comment.GetAnswerId() == "" {
		answerId = nil
	}
	if comment.GetParentId() == "" {
		parentId = nil
	}
	query := "INSERT INTO comment (comment_id, video_id, parent_id, answer_id, user_id, video_comment) VALUES (?, ?, ?, ?, ?, ?)"
	_,err := cr.mysql.ExecContext(ctx, query, comment.GetCommentId(), comment.GetVideoId(),  parentId, answerId, comment.GetUserId(), comment.GetVideoComment())
	if err != nil {
		logger.Error("Error trying ExecContext", err, zap.String("journey", "CreateComment Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	return nil
}