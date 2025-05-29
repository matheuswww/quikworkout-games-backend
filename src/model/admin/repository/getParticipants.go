package admin_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	admin_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/response"
	"go.uber.org/zap"
)


func (ar *adminRepository) GetParticipants(editionID, cursor_createdAt, cursor_userTime string) ([]admin_response.Participant, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if editionID == "" {
		var edition_id string
		query := "SELECT edition_id FROM edition ORDER BY created_at DESC LIMIT 1"
		err := ar.mysql.QueryRowContext(ctx, query).Scan(&edition_id)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipants Repository"))
				return nil,rest_err.NewNotFoundError("no edition found")
			}
			logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipants Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		editionID = edition_id
	}

	var args []any
	args = append(args, editionID)
	query := "SELECT p.video_id, p.placing, e.number, p.user_time, p.desqualified, p.checked, u.user_id, u.name, u.user, p.created_at FROM participant AS p JOIN user_games AS u ON p.user_id = u.user_id JOIN edition AS e ON p.edition_id = e.edition_id WHERE p.edition_id = ? AND "
	if cursor_createdAt != "" {
		query += "p.created_at <= ? AND "
		args = append(args, cursor_createdAt)
	}
	if cursor_userTime != "" {
		query += "p.user_time >= ? AND "
		args = append(args, cursor_userTime)
	}
	query = query[:len(query) - 4]

	query += "ORDER BY p.desqualified ASC, p.user_time IS NOT NULL, p.user_time ASC, p.created_at DESC LIMIT 10"

	rows, err := ar.mysql.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying QueryContext", err, zap.String("journey", "GetParticipants Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()
	var participants []admin_response.Participant
	for rows.Next() {
		var video_id, user_id, name, user, created_at string
		var userTime, placing, desqualified sql.NullString
		var checked bool
		var number int
		err = rows.Scan(&video_id, &placing, &number, &userTime, &desqualified, &checked, &user_id, &name, &user, &created_at)
		if err != nil {
			logger.Error("Error trying Scan", err, zap.String("journey", "GetParticipants Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		var userTimeValid any = nil
		var placingValid any = nil
		var desqualifiedValid any = nil
		if desqualified.Valid {
			desqualifiedValid = desqualified.String
		} 
		if placing.Valid {
			placingValid = placing.String
		}
		if userTime.Valid {
			userTimeValid = userTime.String
		}
		participants = append(participants, admin_response.Participant{
			Video: video_id,
			Edition: number,
			Placing: placingValid,
			UserTime: userTimeValid,
			Desqualified: desqualifiedValid,
			Checked: checked,
			User: admin_response.User{
				UserId: user_id,
				Name: name,
				User: user,
			},
			CreatedAt: created_at,
		})
	}

	if len(participants) == 0 {
		logger.Error("Error trying get participants", errors.New("not found"), zap.String("journey", "GetParticipants Repository"))
		return nil, rest_err.NewNotFoundError("no participants were found")
	}
	
	return participants, nil
}