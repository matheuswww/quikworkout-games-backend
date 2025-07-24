package admin_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	"go.uber.org/zap"
)

func (ar *adminRepository) CreateEdition(createEditionRequest *admin_request.CreateEdition, tops []admin_request.Top, challenges []admin_request.Challenge, savePdf func(id string) *rest_err.RestErr) (*rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var clothing_id string
	query := "SELECT clothing_id FROM clothing WHERE name = ?"
	err := ar.mysql.QueryRowContext(ctx, query, createEditionRequest.ClothingName).Scan(&clothing_id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error clothing not found", errors.New("clothing not found"), zap.String("journey", "CreateEdition Repository"))
			return rest_err.NewBadRequestError("roupa não encontrada")
		}
		logger.Error("Error trying get clothing_id", err, zap.String("journey", "CreateEdition Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	id := uuid.NewString()
	restErr := savePdf(id)
	if restErr != nil {
		return restErr
	}

	tx, err := ar.mysql.Begin()
	if err != nil {
		logger.Error("Error trying init tx", err, zap.String("journey", "CreateEdition Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	query = "INSERT INTO edition (edition_id, start_date, closing_date, clothing_id) VALUES (?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, query, id, createEditionRequest.StartDate, createEditionRequest.ClosingDate, clothing_id)
	if err != nil {
		logger.Error("Error trying insert edition", err, zap.String("journey", "CreateEdition Repository"))
		err = tx.Rollback()
		if err != nil {
			logger.Error("Error trying rollback", err, zap.String("journey", "CreateEdition Repository"))
			return rest_err.NewInternalServerError("server error")
		}
		return rest_err.NewInternalServerError("server error")
	}
	
	for i := 0; i < len(tops); i++ {
		query = "INSERT INTO top (edition_id, top, gain, category) VALUES (?, ?, ?, ?)"
		_, err := tx.ExecContext(ctx, query, id, tops[i].Top, tops[i].Gain, tops[i].Category)
		if err != nil {
			logger.Error("Error trying insert top", err, zap.String("journey", "CreateEdition Repository"))
			err = tx.Rollback()
			if err != nil {
				logger.Error("Error trying rollback", err, zap.String("journey", "CreateEdition Repository"))
				return rest_err.NewInternalServerError("server error")
			}
			return rest_err.NewInternalServerError("server error")
		}
	}

	for i := 0; i < len(challenges); i++ {
		query = "INSERT INTO challenge (edition_id, challenge, category, sex) VALUES (?, ?, ?, ?)"
		_, err := tx.ExecContext(ctx, query, id, challenges[i].Challenge, challenges[i].Category, challenges[i].Sex)
		if err != nil {
			logger.Error("Error trying insert challenge", err, zap.String("journey", "CreateEdition Repository"))
			err = tx.Rollback()
			if err != nil {
				logger.Error("Error trying rollback", err, zap.String("journey", "CreateEdition Repository"))
				return rest_err.NewInternalServerError("server error")
			}
			return rest_err.NewInternalServerError("server error")
		}
	}
	
	err = tx.Commit()
	if err != nil {
		logger.Error("Error trying commit", err, zap.String("journey", "CreateEdition Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}