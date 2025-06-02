package user_repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	user_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/response"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	"go.uber.org/zap"
)

func (ur *userRepository) GetParticipations(user_domain user_domain.UserDomainInterface, cursor string) (*user_response.GetParticipations, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var name, user string
	query := "SELECT name, user FROM user_games WHERE user_id = ?"
	err := ur.mysql.QueryRowContext(ctx, query, user_domain.GetId()).Scan(&name, &user)
	if err != nil { 
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipantions Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}

	query = "SELECT p.video_id, p.placing, e.number, p.user_time, p.desqualified, p.checked, p.created_at, t.gain FROM participant AS p JOIN edition AS e ON p.edition_id = e.edition_id LEFT JOIN top AS t ON p.placing = t.top AND t.edition_id = p.edition_id AND p.placing IS NOT NULL WHERE p.user_id = ? "
	var args []any
	args = append(args, user_domain.GetId())
	if cursor != "" {
		query += "AND p.created_at < ? "
		args = append(args, cursor)
	}
	query += "ORDER BY created_at DESC LIMIT 10"
	
	rows, err := ur.mysql.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipantions Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}

	var participants []user_response.Participantion
	for rows.Next() {
		var video_id, created_at string
		var number int
		var gain sql.NullInt64
		var checked bool
		var placing, user_time, desqualified sql.NullString
		err := rows.Scan(&video_id, &placing, &number, &user_time, &desqualified, &checked, &created_at, &gain)
		if err != nil {
			logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipantions Repository"))
			return nil, rest_err.NewInternalServerError("server error")
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
			Video: video_id,
			Placing: validPlacing,
			Edition: number,
			Gain: validGain,
			UserTime: validUserTime,
			Desqualified: validDesqualified,
			Checked: checked,
			CreatedAt: created_at,
		})
	}

	if len(participants) == 0 {
		return nil, rest_err.NewNotFoundError("no participation was found")
	}
	return &user_response.GetParticipations{
		Participations: participants,
		User: user_response.User{
			UserId: user_domain.GetId(),
			Name: name,
			User: user,
		},
	}, nil
}