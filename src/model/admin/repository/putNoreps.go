package admin_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	"go.uber.org/zap"
)

func (ar *adminRepository) PutNoreps(putNorepsRequest *admin_request.PutNoreps) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	val, restErr := ar.HasPlacing(putNorepsRequest.EditionId, putNorepsRequest.Category, putNorepsRequest.Sex)
	if restErr != nil {
		return restErr
	}
	if val {
		return rest_err.NewBadRequestError("this edition was finished")
	}

	var desqualified sql.NullString
	query := "SELECT desqualified FROM participant WHERE video_id = ?"
	err := ar.mysql.QueryRowContext(ctx, query, putNorepsRequest.VideoId).Scan(&desqualified)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error video not found", errors.New("video not found"), zap.String("journey", "PutTime Repository"))
			return rest_err.NewBadRequestError("video not found")
		}
		logger.Error("Error trying get participant", err, zap.String("journey", "PutTime Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	if desqualified.Valid {
		logger.Error("Error trying PutTime", errors.New("this user is desqualified"), zap.String("journey", "PutTime Repository"))
		return rest_err.NewBadRequestError("this user is desqualified")
	}

	var noreps string
	for _, norep := range putNorepsRequest.Noreps {
		noreps += fmt.Sprintf("%s - %s\n", norep.Time, norep.NoRep)
	}

	query = "UPDATE participant SET noreps = ? WHERE video_id = ?"
	_,err = ar.mysql.ExecContext(ctx, query, noreps, putNorepsRequest.VideoId)
	if err != nil {
		logger.Error("Error trying update noreps", err, zap.String("journey", "PutNoreps Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	return nil
}