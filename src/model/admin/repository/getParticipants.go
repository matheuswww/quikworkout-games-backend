package admin_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	admin_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/response"
	"go.uber.org/zap"
)

func (ar *adminRepository) GetParticipants(getParticipantsRequest *admin_request.GetParticipants) (*admin_response.GetParticipants, *sql.DB, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if getParticipantsRequest.Category == "" && getParticipantsRequest.VideoId == "" {
		logger.Error("Error trying get participants", errors.New("category or video_id must be provided"), zap.String("journey", "GetParticipant Repository"))
		return nil, nil, rest_err.NewBadRequestError("category or video_id must be provided")
	}
	var closing_date string
	var edition_id string
	var query string
	var args []any
	if getParticipantsRequest.EditionId == "" {
		query = "SELECT edition_id, closing_date FROM edition ORDER BY created_at DESC LIMIT 1"
	} else {
		query = "SELECT edition_id, closing_date FROM edition WHERE edition_id = ?"
		args = append(args, getParticipantsRequest.EditionId)
	}

	err := ar.mysql.QueryRowContext(ctx, query, args...).Scan(&edition_id, &closing_date)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipants Repository"))
			return nil, nil, rest_err.NewNotFoundError("no edition found")
		}
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipants Repository"))
		return nil, nil, rest_err.NewInternalServerError("server error")
	}
	getParticipantsRequest.EditionId = edition_id
	args = nil

	args = append(args, getParticipantsRequest.EditionId)
	moreDataArgs := []any{getParticipantsRequest.EditionId}
	from := "FROM participant AS p JOIN user_games AS u ON p.user_id = u.user_id JOIN user AS uq ON uq.user_id = u.user_id JOIN edition AS e ON p.edition_id = e.edition_id LEFT JOIN top AS t ON t.top = p.placing AND t.edition_id = p.edition_id AND t.category = p.category JOIN challenge AS c ON c.edition_id = p.edition_id AND c.category = p.category WHERE p.edition_id = ? AND "
	moreDataQuery := "SELECT 1 " + from
	query = "SELECT p.video_id, p.placing, p.edition_id, e.number, p.user_time, p.desqualified, p.category, c.challenge, p.sent, p.checked, u.user_id, u.name, u.user, t.gain, uq.email, p.created_at " + from
	if getParticipantsRequest.Category != "" {
		query += "p.category = ? AND "
		moreDataQuery += "p.category = ? AND "
		args = append(args, getParticipantsRequest.Category)
		moreDataArgs = append(moreDataArgs, getParticipantsRequest.Category)
	}
	if getParticipantsRequest.VideoId != "" {
		query += "p.video_id = ? AND "
		moreDataQuery += "p.video_id = ? AND "
		args = append(args, getParticipantsRequest.VideoId)
		moreDataArgs = append(moreDataArgs, getParticipantsRequest.VideoId)
	}
	if getParticipantsRequest.CursorCreatedAt != "" {
		query += "(p.created_at < ? OR p.user_time IS NOT NULL) AND "
		args = append(args, getParticipantsRequest.CursorCreatedAt)
	}
	if getParticipantsRequest.CursorUserTime != "" {
		query += "p.user_time > ? AND "
		args = append(args, getParticipantsRequest.CursorUserTime)
	}
	query = query[:len(query)-4]

	order := "ORDER BY p.desqualified ASC, p.placing, p.placing ASC, p.user_time IS NOT NULL, p.user_time ASC, p.created_at DESC "
	query += order+"LIMIT 10"

	rows, err := ar.mysql.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying QueryContext", err, zap.String("journey", "GetParticipants Repository"))
		return nil, nil, rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()
	var participants []admin_response.Participant
	for rows.Next() {
		var video_id, user_id, edition_id, name, category, challenge, user, email, created_at string
		var userTime, placing, desqualified sql.NullString
		var gain sql.NullInt64
		var checked, sent bool
		var number int
		err = rows.Scan(&video_id, &placing, &edition_id, &number, &userTime, &desqualified, &category, &challenge, &sent, &checked, &user_id, &name, &user, &gain, &email, &created_at)
		if err != nil {
			logger.Error("Error trying Scan", err, zap.String("journey", "GetParticipants Repository"))
			return nil, nil, rest_err.NewInternalServerError("server error")
		}
		var gainValid any = nil
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
		if gain.Valid {
			gainValid = gain.Int64
		}
		participants = append(participants, admin_response.Participant{
			VideoId:      video_id,
			Edition:      number,
			EditionId:    edition_id,
			Sent:         sent,
			Placing:      placingValid,
			Category:     category,
			Challenge:    challenge,
			Gain:         gainValid,
			UserTime:     userTimeValid,
			Desqualified: desqualifiedValid,
			Checked:      checked,
			User: admin_response.User{
				UserId: user_id,
				Name:   name,
				User:   user,
				Email:  email,
			},
			CreatedAt: created_at,
		})
	}

	if len(participants) == 0 {
		logger.Error("Error trying get participants", errors.New("not found"), zap.String("journey", "GetParticipants Repository"))
		return nil, nil, rest_err.NewNotFoundError("no participants were found")
	}

	last := participants[len(participants)-1]
	moreDataQuery += "(p.created_at < ? OR p.user_time IS NOT NULL) AND "
	moreDataArgs = append(moreDataArgs, last.CreatedAt)
	if last.UserTime != nil {
		moreDataQuery += "p.user_time > ? AND "
		moreDataArgs = append(moreDataArgs, last.UserTime)
	}
	moreDataQuery = moreDataQuery[:len(moreDataQuery)-4]
	moreDataQuery += order

	more := false
	err = ar.mysql.QueryRowContext(ctx, moreDataQuery, moreDataArgs...).Scan(&more)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error trying QueryRowContext", err, zap.String("journey", "GetParticipants Repository"))
		return nil, nil, rest_err.NewInternalServerError("server error")
	}

	return &admin_response.GetParticipants{
		Participants: participants,
		ClosingDate:  closing_date,
		More:         more,
	}, ar.mysql, nil
}
