package user_repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/request"
	user_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/response"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	"go.uber.org/zap"
)

func (ur *userRepository) GetParticipations(user_domain user_domain.UserDomainInterface, getParticipationsRequest *user_request.GetParticipations) (*user_response.GetParticipations, *sql.DB, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var name, user string
	query := "SELECT name, user FROM user_games WHERE user_id = ?"
	err := ur.mysql.QueryRowContext(ctx, query, user_domain.GetId()).Scan(&name, &user)
	if err != nil { 
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipantions Repository"))
		return nil, nil, rest_err.NewInternalServerError("server error")
	}

	query = "SELECT p.video_id, p.placing, p.edition_id, e.number, p.user_time, p.desqualified, p.category, p.sent, p.checked, p.created_at, t.gain FROM participant AS p JOIN edition AS e ON p.edition_id = e.edition_id LEFT JOIN top AS t ON p.placing = t.top AND t.edition_id = p.edition_id AND t.category = p.category AND p.placing IS NOT NULL WHERE p.user_id = ? AND "
	var args []any
	args = append(args, user_domain.GetId())
	if getParticipationsRequest.VideoId != "" {
		query += "p.video_id = ? AND "
		args = append(args, getParticipationsRequest.VideoId)
	}
	if getParticipationsRequest.EditionId != "" {
		query += "p.edition_id = ? AND "
		args = append(args, getParticipationsRequest.EditionId)
	}
	if getParticipationsRequest.Cursor != "" {
		query += "p.created_at < ? AND "
		args = append(args, getParticipationsRequest.Cursor)
	}
	if len(args) > 0 {
		query = query[:len(query) - 4]
	}
	query += "ORDER BY created_at DESC LIMIT ?"
	if getParticipationsRequest.Limit > 10 || getParticipationsRequest.Limit == 0 {
		getParticipationsRequest.Limit = 10
	}
	args = append(args, getParticipationsRequest.Limit)
	rows, err := ur.mysql.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipantions Repository"))
		return nil, nil, rest_err.NewInternalServerError("server error")
	}

	var participants []user_response.Participantion
	for rows.Next() {
		var video_id, edition_id, category, created_at string
		var number int
		var gain sql.NullInt64
		var checked, sent bool
		var placing, user_time, desqualified sql.NullString
		err := rows.Scan(&video_id, &placing, &edition_id, &number, &user_time, &desqualified, &category, &sent, &checked, &created_at, &gain)
		if err != nil {
			logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipantions Repository"))
			return nil, nil, rest_err.NewInternalServerError("server error")
		}
		var validPlacing any = nil
		var validUserTime any = nil
		var validDesqualified any = nil
		var validGain any = nil
		if gain.Valid {
			validGain = gain.Int64
		}
		if desqualified.Valid {
			validDesqualified = desqualified.String
		}
		if placing.Valid {
			validPlacing = placing.String
		}
		if user_time.Valid {
			validUserTime = user_time.String
		}
		participants = append(participants, user_response.Participantion{
			VideoId: video_id,
			Placing: validPlacing,
			Edition: number,
			EditionId: edition_id,
			Sent: sent,
			Gain: validGain,
			UserTime: validUserTime,
			Category: category,
			Desqualified: validDesqualified,
			Checked: checked,
			CreatedAt: created_at,
		})
	}

	if len(participants) == 0 {
		return nil, nil, rest_err.NewNotFoundError("no participation was found")
	}
	return &user_response.GetParticipations{
		Participations: participants,
		User: user_response.User{
			UserId: user_domain.GetId(),
			Name: name,
			User: user,
		},
	}, ur.mysql, nil
}