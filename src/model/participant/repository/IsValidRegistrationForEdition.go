package participant_repository

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/vimeo"
	participant_domain "github.com/matheuswww/quikworkout-games-backend/src/model/participant"
	"go.uber.org/zap"
)

func (pr *participantRepository) IsValidRegistrationForEdition(participantDomain participant_domain.ParticipantDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	logger.Info("Init IsValidRegistrationForEdition Repository", zap.String("journey", "IsValidRegistrationForEdition Repository"))

	var edition_id, start_date, closing_date string
	query := "SELECT edition_id, start_date, closing_date FROM edition ORDER BY created_at DESC LIMIT 1"
	err := pr.mysql.QueryRowContext(ctx, query).Scan(&edition_id, &start_date, &closing_date)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error trying QueryRowContext", err, zap.String("journey", "IsValidRegistrationForEdition Repository"))
			return rest_err.NewBadRequestError("no edition found")
		}
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "IsValidRegistrationForEdition Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	participantDomain.SetEditionID(edition_id)

	format := "2006-01-02"
	start_date_formated, err1 := time.Parse(format, start_date)
	closing_date_formated, err2 := time.Parse(format, closing_date)
	if err1 != nil || err2 != nil {
		logger.Error("Error trying Parse date", errors.New(err1.Error()+" "+err2.Error()), zap.String("journey", "IsValidRegistrationForEdition Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	closing_date_formated = closing_date_formated.Add(24*time.Hour - time.Second)
	now := time.Now()
	if now.Before(start_date_formated) || now.After(closing_date_formated) {
		logger.Error("Error trying IsValidRegistrationForEdition", errors.New("it is no longer possible to register"), zap.String("journey", "IsValidRegistrationForEdition Repository"))
		return rest_err.NewBadRequestError("is not possible to register")
	}

	var video_id string
	query = "SELECT video_id FROM participant WHERE user_id = ? AND edition_id = ?"
	err = pr.mysql.QueryRowContext(ctx, query, participantDomain.GetUserID(), participantDomain.GetEditionID()).Scan(&video_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "IsValidRegistrationForEdition Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	_, status, err := vimeo.GetVideo(vimeo.GetVideoParams{ VideoID: video_id })
	if (status != http.StatusOK && status != http.StatusNotFound) {
		logger.Error("Error trying GetVideo", err, zap.String("journey", "IsValidRegistrationForEdition Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	if status != http.StatusNotFound {
		logger.Error("Error trying IsValidRegistrationForEdition", errors.New("user is already in editing"), zap.String("journey", "IsValidRegistrationForEdition Repository"))
		return rest_err.NewBadRequestError("user is already in edition")
	}
	query = "DELETE FROM participant WHERE user_id = ? AND edition_id = ? AND video_id = ?"
	_,err = pr.mysql.ExecContext(ctx, query, participantDomain.GetUserID(), participantDomain.GetEditionID(), video_id)
	if err != nil {
		logger.Error("Error trying ExecContext", err, zap.String("journey", "IsValidRegistrationForEdition Repository"))
		return rest_err.NewInternalServerError("server error")
	}
	return nil
}

