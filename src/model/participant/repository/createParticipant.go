package participant_repository

import (
	"context"

	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	participant_domain "github.com/matheuswww/quikworkout-games-backend/src/model/participant"
	"go.uber.org/zap"
)

func (cr *participantRepository) CreateParticipant(participantDomain participant_domain.ParticipantDomainInterface, instagram string) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	tx, err := cr.mysql.Begin()
	if err != nil {
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "CreateParticiapnt Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	query := "INSERT INTO participant (video_id, user_id, edition_id, sent, checked) VALUES (?, ?, ?, ?, ?)"
	_,err = tx.ExecContext(ctx, query, participantDomain.GetVideoID(), participantDomain.GetUserID(), participantDomain.GetEditionID(), participantDomain.GetSent(), participantDomain.GetChecked())
	if err != nil {
		logger.Error("Error trying ExecContext", err, zap.String("journey", "CreateParticiapnt Repository"))
		err := tx.Rollback()
		if err != nil {
			logger.Error("Error trying rollback", err, zap.String("journey", "CreateParticipant Repository"))
		}
		return rest_err.NewInternalServerError("server error")
	}

	query = "UPDATE user_games SET instagram = ? WHERE user_id = ?"
	_,err = tx.ExecContext(ctx, query, instagram, participantDomain.GetUserID())
	if err != nil {
		logger.Error("Error trying ExecContext", err, zap.String("journey", "CreateParticiapnt Repository"))
		err := tx.Rollback()
		if err != nil {
			logger.Error("Error trying rollback", err, zap.String("journey", "CreateParticipant Repository"))
		}
		return rest_err.NewInternalServerError("server error")
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Error trying commit", err, zap.String("journey", "CreateParticiapnt Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	return nil
}