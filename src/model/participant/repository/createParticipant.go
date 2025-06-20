package participant_repository

import (
	"context"

	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	participant_domain "github.com/matheuswww/quikworkout-games-backend/src/model/participant"
	"go.uber.org/zap"
)

func (cr *participantRepository) CreateParticipant(participantDomain participant_domain.ParticipantDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := "INSERT INTO participant (video_id, user_id, edition_id, sent, checked) VALUES (?, ?, ?, ?, ?)"
	_,err := cr.mysql.ExecContext(ctx, query, participantDomain.GetVideoID(), participantDomain.GetUserID(), participantDomain.GetEditionID(), participantDomain.GetSent(), participantDomain.GetChecked())
	if err != nil {
		logger.Error("Error trying ExecContext", err, zap.String("journey", "CreateParticiapnt Repository"))
		return rest_err.NewInternalServerError("server error")
	}

	return nil
}