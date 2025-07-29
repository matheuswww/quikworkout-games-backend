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

	from := "FROM participant AS p JOIN edition AS e ON p.edition_id = e.edition_id LEFT JOIN top AS t ON p.placing = t.top AND t.edition_id = p.edition_id AND t.category = p.category AND p.placing IS NOT NULL WHERE p.user_id = ? AND "
	moreDataQuery := "SELECT 1 " + from
	query = "SELECT p.video_id, p.placing, p.edition_id, e.number, p.user_time, p.final_time, p.desqualified, p.category, p.noreps, p.sex, p.sent, p.checked, p.created_at, t.gain " + from
	var args []any
	var moreDataArgs []any
	args = append(args, user_domain.GetId())
	moreDataArgs = append(moreDataArgs, user_domain.GetId())
	if getParticipationsRequest.VideoId != "" {
		query += "p.video_id = ? AND "
		moreDataQuery += "p.video_id = ? AND "
		args = append(args, getParticipationsRequest.VideoId)
		moreDataArgs = append(moreDataArgs, getParticipationsRequest.VideoId)
	}
	if getParticipationsRequest.EditionId != "" {
		query += "p.edition_id = ? AND "
		moreDataQuery += "p.edition_id = ? AND "
		args = append(args, getParticipationsRequest.EditionId)
		moreDataArgs = append(moreDataArgs, getParticipationsRequest.EditionId)
	}
	if getParticipationsRequest.Cursor != "" {
		query += "p.created_at < ? AND "
		moreDataQuery += "p.created_at < ? AND "
		args = append(args, getParticipationsRequest.Cursor)
		moreDataArgs = append(moreDataArgs, getParticipationsRequest.Cursor)
	}
	if len(args) > 0 {
		query = query[:len(query)-4]
	}
	order := "ORDER BY p.created_at DESC "
	query += order + "LIMIT ?"
	if getParticipationsRequest.Limit > 10 || getParticipationsRequest.Limit == 0 {
		getParticipationsRequest.Limit = 10
	}
	args = append(args, getParticipationsRequest.Limit)
	rows, err := ur.mysql.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying get participantions", err, zap.String("journey", "GetParticipantions Repository"))
		return nil, nil, rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()

	var participants []user_response.Participantion
	for rows.Next() {
		var video_id, edition_id, category, user_time, sex, created_at string
		var number int
		var gain sql.NullInt64
		var checked, sent bool
		var placing, final_time, desqualified, noreps sql.NullString
		err := rows.Scan(&video_id, &placing, &edition_id, &number, &user_time, &final_time, &desqualified, &category, &noreps, &sex, &sent, &checked, &created_at, &gain)
		if err != nil {
			logger.Error("Error trying scan", err, zap.String("journey", "GetParticipantions Repository"))
			return nil, nil, rest_err.NewInternalServerError("server error")
		}
		var validPlacing any = nil
		var validFinalTime any = nil
		var validDesqualified any = nil
		var validGain any = nil
		var validNoreps any = nil

		if noreps.Valid {
			validNoreps = noreps.String
		}
		if gain.Valid {
			validGain = gain.Int64
		}
		if desqualified.Valid {
			validDesqualified = desqualified.String
		}
		if placing.Valid {
			validPlacing = placing.String
		}
		if final_time.Valid {
			validFinalTime = final_time.String
		}
		participants = append(participants, user_response.Participantion{
			VideoId:      video_id,
			Placing:      validPlacing,
			Edition:      number,
			EditionId:    edition_id,
			Sent:         sent,
			Gain:         validGain,
			FinalTime:    validFinalTime,
			UserTime:     user_time,
			Category:     category,
			Noreps:       validNoreps,
			Sex:          sex,
			Desqualified: validDesqualified,
			Checked:      checked,
			CreatedAt:    created_at,
		})
	}

	if len(participants) == 0 {
		return nil, nil, rest_err.NewNotFoundError("no participation was found")
	}

	last := participants[len(participants)-1]
	moreDataQuery += "p.created_at < ? "
	moreDataArgs = append(moreDataArgs, last.CreatedAt)

	moreDataQuery += order + "LIMIT 1"

	more := false
	err = ur.mysql.QueryRowContext(ctx, moreDataQuery, moreDataArgs...).Scan(&more)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error trying get participations", err, zap.String("journey", "GetParticipantions Repository"))
		return nil, nil, rest_err.NewInternalServerError("server error")
	}

	return &user_response.GetParticipations{
		Participations: participants,
		User: user_response.User{
			UserId: user_domain.GetId(),
			Name:   name,
			User:   user,
		},
		More: more,
	}, ur.mysql, nil
}
