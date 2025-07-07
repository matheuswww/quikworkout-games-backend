package admin_repository

import (
	"context"
	"errors"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ar *adminRepository) PutTime(videoId, editionId, userTime string) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	val, restErr := ar.HasNoPlacing(editionId)
	if restErr != nil {
		return restErr
	}
	if !val {
		return rest_err.NewBadRequestError("this edition was finished")
	}

	query := "SELECT COUNT(*) FROM participant WHERE video_id = ?"
	var count int
	err := ar.mysql.QueryRowContext(ctx, query, videoId).Scan(&count)
	if err != nil {
		logger.Error("Error trying get participant", err, zap.String("journey", "PutTime Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if count == 0 {
		logger.Error("Error video not found", errors.New("video not found"), zap.String("journey", "PutTime Repository"))
		return rest_err.NewBadRequestError("video not found")
	}

	var putTimeParam any = userTime
	if userTime == "" {
		putTimeParam = nil		
	}

	query = "UPDATE participant SET user_time = ? WHERE video_id = ?"
	_,err = ar.mysql.ExecContext(ctx, query, putTimeParam, videoId)
	if err != nil {
		logger.Error("Error trying update participant", err, zap.String("journey", "PutTime Repository"))
		return rest_err.NewInternalServerError("server Error")
	}

	return nil
}