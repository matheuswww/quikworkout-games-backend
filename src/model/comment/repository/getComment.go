package comment_repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	comment_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/comment/response"
	"go.uber.org/zap"
)

func (cr *commentRepository) GetComment(video_id, cursor, commentId string) ([]comment_response.GetComment, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	
	var query string
	var args []any
	args = append(args, video_id)
	if commentId == "" {
		query = "SELECT c.comment_id, c.parent_id, c.answer_id, c.video_comment, c.user_id, u.name, u.user, u.category, c.created_at FROM comment AS c JOIN user_games AS u ON u.user_id = c.user_id WHERE c.video_id = ? AND c.parent_id IS NULL "
		if cursor != "" {
			query += "AND c.created_at > ? "
		}
		query += "ORDER BY c.created_at ASC LIMIT 10"
	} else {
		query = "SELECT c.comment_id, c.parent_id, c.answer_id, c.video_comment, c.user_id, u.name, u.user, u.category, c.created_at FROM comment AS c JOIN user_games AS u ON u.user_id = c.user_id WHERE video_id = ? AND c.parent_id = ? "
		args = append(args, commentId)
		if cursor != "" {
			query += "AND c.created_at > ? "
		}
		query += "ORDER BY c.created_at ASC LIMIT 5"
	}
	if cursor != "" {
		args = append(args, cursor)
	}

	rows, err := cr.mysql.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying QueryContext", err, zap.String("journey", "GetComment Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}

	var comments []comment_response.GetComment
	for rows.Next() {
		var commentId, comment, userId, name, user, category, created_at string
		var parentId, answerId sql.NullString
		err := rows.Scan(&commentId, &parentId, &answerId, &comment, &userId, &name, &user, &category, &created_at)
		if err != nil {
			logger.Error("Error trying QueryContext", err, zap.String("journey", "GetComment Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		var validParentId any = nil
		var validAnswerId any = nil
		if answerId.Valid {
			validAnswerId = answerId.String
		}
		if parentId.Valid {
			validParentId = parentId.String
		}
		comments = append(comments, comment_response.GetComment{
			CommentId: commentId,
			ParentId:  validParentId,
			Answer:    validAnswerId,
			Comment:   comment,
			UserId:    userId,
			Name:      name,
			User:      user,
			Category:  category,
			CreatedAt: created_at,
		})
	}

	return comments, nil
}