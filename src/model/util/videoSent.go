package model_util

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

func VideoSent(db *sql.DB, videoId, userId string) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	fmt.Println("success")
	query := "SELECT COUNT(*) FROM participant WHERE video_id = ? AND user_id = ?"
	var count int
	err := db.QueryRowContext(ctx, query, videoId, userId).Scan(&count)
	if err != nil {
		logger.Error("Error trying get participant", err, zap.String("journey", "Video Sent"))
		return rest_err.NewInternalServerError("server error")
	}
	if count == 0 {
		return rest_err.NewBadRequestError("video not found")
	}
	query = "UPDATE participant SET sent = TRUE WHERE video_id = ? AND user_id = ?"
	_,err = db.ExecContext(ctx, query, videoId, userId)
	if err != nil {
		logger.Error("Error trying update participant", err, zap.String("journey", "Video Sent"))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}