package participant_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	participant_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/request"
	participant_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/response"
	"go.uber.org/zap"
)

func (pr *participantRepository) GetParticipants(getParticipantRequest *participant_request.GetParticipant) ([]participant_response.Participant, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if getParticipantRequest.EditionId == "" {
		var edition_id string
		query := "SELECT edition_id FROM edition ORDER BY created_at DESC LIMIT 1"
		err := pr.mysql.QueryRowContext(ctx, query).Scan(&edition_id)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipant Repository"))
				return nil,rest_err.NewNotFoundError("no edition found")
			}
			logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipant Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		getParticipantRequest.EditionId = edition_id
	}

	var args []any
	args = append(args, getParticipantRequest.EditionId)
	query := "SELECT p.video_id, u.user_id, u.name, u.user, p.edition_id, p.user_time, p.placing, p.created_at FROM participant AS p JOIN user_games AS u ON p.user_id = u.user_id WHERE p.edition_id = ? AND p.checked IS true AND p.sent IS true AND desqualified IS NULL AND "
	if getParticipantRequest.NotVideoId != "" {
		query += "p.video_id != ? AND "
		args = append(args, getParticipantRequest.NotVideoId)
	}
	if getParticipantRequest.VideoId != "" {
		query += "p.video_id = ? AND "
		args = append(args, getParticipantRequest.VideoId)
	}
	if getParticipantRequest.CursorCreatedAt != "" {
		query += "p.created_at < ? AND "
		args = append(args, getParticipantRequest.CursorCreatedAt)
	}
	if getParticipantRequest.CursorUserTime != "" {
		signal := "> ?"
		if getParticipantRequest.WorstTime {
			signal = "< ?"
		}
		query += "(p.user_time "+signal+" OR p.user_time IS NULL) AND "
		args = append(args, getParticipantRequest.CursorUserTime)
	}
	query = query[:len(query) - 4]

	query += "ORDER BY "
	query += "p.user_time IS NULL, "
	if !getParticipantRequest.WorstTime {
		query += "p.user_time ASC, "
	} else if getParticipantRequest.WorstTime {
		query += "p.user_time DESC, "
	}
	query += "p.created_at DESC "
	query += "LIMIT 10"

	rows, err := pr.mysql.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying QueryContext", err, zap.String("journey", "GetParticipant Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()
	var participants []participant_response.Participant
	for rows.Next() {
		var video_id, user_id, name, user, edition_id, created_at string
		var userTime, placing sql.NullString
		err = rows.Scan(&video_id, &user_id, &name, &user, &edition_id, &userTime, &placing, &created_at)
		if err != nil {
			logger.Error("Error trying Scan", err, zap.String("journey", "GetParticipant Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		var userTimeValid any = nil
		var userPlacingValid any = nil
		if placing.Valid {
			userPlacingValid = placing.String
		}
 		if userTime.Valid {
			userTimeValid = userTime.String
		}
		participants = append(participants, participant_response.Participant{
			VideoId: video_id,
			UserTime: userTimeValid,
			Edition_id: edition_id,
			Placing: userPlacingValid,
			User: participant_response.User{
				UserId: user_id,
				Name: name,
				User: user,
			},
			CreatedAt: created_at,
		})
	}

	if len(participants) == 0 {
		logger.Error("Error trying get participants", errors.New("not found"), zap.String("journey", "GetParticipant Repository"))
		return nil, rest_err.NewNotFoundError("no participants were found")
	}
	
	return participants, nil
}