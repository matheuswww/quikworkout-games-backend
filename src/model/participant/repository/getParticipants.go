package participant_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	participant_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/response"
	"go.uber.org/zap"
)

func (pr *participantRepository) GetParticipants(editionID, cursor_createdAt, cursor_userTime string, worstTime bool) ([]participant_response.Participant, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if editionID == "" {
		var edition_id string
		query := "SELECT edition_id FROM edition ORDER BY created_at DESC LIMIT 1"
		err := pr.mysql.QueryRowContext(ctx, query).Scan(&edition_id)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipant Repository"))
				return nil,rest_err.NewBadRequestError("no edition found")
			}
			logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipant Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		editionID = edition_id
	}

	var args []any
	args = append(args, editionID)
	query := "SELECT p.video_id, u.user_id, u.name, u.user, p.user_time, p.created_at FROM participant AS p JOIN user_games AS u ON p.user_id = u.user_id WHERE p.edition_id = ? AND p.checked IS true AND desqualified IS NULL AND "
	if cursor_createdAt != "" {
		query += "p.created_at <= ? AND "
		args = append(args, cursor_createdAt)
	}
	if cursor_userTime != "" {
		signal := ">= ?"
		if worstTime {
			signal = "<= ?"
		}
		query += "(p.user_time "+signal+" OR p.user_time IS NULL) AND "
		args = append(args, cursor_userTime)
	}
	query = query[:len(query) - 4]

	query += "ORDER BY "
	query += "p.user_time IS NULL, "
	if !worstTime {
		query += "p.user_time ASC, "
	} else if worstTime {
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
		var video_id, user_id, name, user, created_at string
		var userTime sql.NullString
		err = rows.Scan(&video_id, &user_id, &name, &user, &userTime, &created_at)
		if err != nil {
			logger.Error("Error trying Scan", err, zap.String("journey", "GetParticipant Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		var userTimeValid any = nil
		if userTime.Valid {
			userTimeValid = userTime.String
		}
		participants = append(participants, participant_response.Participant{
			Video: video_id,
			UserTime: userTimeValid,
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
		return nil, rest_err.NewBadRequestError("no participants were found")
	}
	
	return participants, nil
}