package admin_repository

import (
	"context"
	"errors"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ar *adminRepository) DesqualifyVideo(videoID, editionId, category, sex, desqualifed string) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	val, restErr := ar.HasPlacing(editionId, category, sex)
	if restErr != nil {
		return restErr
	}
	if val {
		return rest_err.NewBadRequestError("this edition was finished")
	}

	query := "SELECT COUNT(*) FROM participant WHERE video_id = ?"
	var count int
	err := ar.mysql.QueryRowContext(ctx, query, videoID).Scan(&count)
	if err != nil {
		logger.Error("Error trying get count", err, zap.String("journey", "DesqualifyVideo Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if count == 0 {
		logger.Error("Error video not found", errors.New("video not found"), zap.String("journey", "DesqualifyVideo Repository"))
		return rest_err.NewBadRequestError("video not found")
	}

	var desqualifedParam any = desqualifed
	if desqualifed == "" {
		desqualifedParam = nil		
	}
	query = "UPDATE participant SET desqualified = ?, final_time = NULL WHERE video_id = ?"
	_, err = ar.mysql.ExecContext(ctx, query, desqualifedParam, videoID)
	if err != nil {
		logger.Error("Error trying update participant", err, zap.String("journey", "DesqualifyVideo Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	return nil
}