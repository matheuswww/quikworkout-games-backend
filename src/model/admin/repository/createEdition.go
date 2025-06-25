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

func (ar *adminRepository) CreateEdition(createEditionRequest *admin_request.CreateEdition) *rest_err.RestErr {
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
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "CreateEdition Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	tx, err := ar.mysql.Begin()
	if err != nil {
		logger.Error("Error trying init tx", err, zap.String("journey", "CreateEdition Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	id := uuid.NewString()
	query = "INSERT INTO edition (edition_id, start_date, closing_date, rules, clothing_id) VALUES (?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, query, id, createEditionRequest.StartDate, createEditionRequest.ClosingDate, createEditionRequest.Rules, clothing_id)
	if err != nil {
		logger.Error("Error trying ExecContext", err, zap.String("journey", "CreateEdition Repository"))
		err = tx.Rollback()
		if err != nil {
			logger.Error("Error trying rollback", err, zap.String("journey", "CreateEdition Repository"))
			return rest_err.NewInternalServerError("server error")
		}
		return rest_err.NewInternalServerError("server error")
	}

	for i := 0; i < len(createEditionRequest.Tops); i++ {
		query = "INSERT INTO top (edition_id, top, gain, category) VALUES (?, ?, ?, ?)"
		_, err := tx.ExecContext(ctx, query, id, createEditionRequest.Tops[i].Top, createEditionRequest.Tops[i].Gain, createEditionRequest.Tops[i].Category)
		if err != nil {
			logger.Error("Error trying ExecContext", err, zap.String("journey", "CreateEdition Repository"))
			err = tx.Rollback()
			if err != nil {
				logger.Error("Error trying rollback", err, zap.String("journey", "CreateEdition Repository"))
				return rest_err.NewInternalServerError("server error")
			}
			return rest_err.NewInternalServerError("server error")
		}
	}

	for i := 0; i < len(createEditionRequest.Challenge); i++ {
		query = "INSERT INTO challenge (edition_id, challenge, category) VALUES (?, ?, ?)"
		_, err := tx.ExecContext(ctx, query, id, createEditionRequest.Challenge[i].Challenge, createEditionRequest.Challenge[i].Category)
		if err != nil {
			logger.Error("Error trying ExecContext", err, zap.String("journey", "CreateEdition Repository"))
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