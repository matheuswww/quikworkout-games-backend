package participant_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	participant_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/request"
	participant_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/response"
	"go.uber.org/zap"
)

var monthNames = []string{
	"janeiro", "fevereiro", "março", "abril", "maio", "junho",
	"julho", "agosto", "setembro", "outubro", "novembro", "dezembro",
}

func (pr *participantRepository) GetParticipants(getParticipantRequest *participant_request.GetParticipant) (*participant_response.GetParticipant, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var closing_date string
	var query string
	var args []any
	if getParticipantRequest.EditionId == "" {
		query = "SELECT edition_id, closing_date FROM edition ORDER BY created_at DESC LIMIT 1"
	} else {
		query = "SELECT edition_id, closing_date FROM edition WHERE edition_id = ?"
		args = append(args, getParticipantRequest.EditionId)
	}

	var edition_id string
	err := pr.mysql.QueryRowContext(ctx, query, args...).Scan(&edition_id, &closing_date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("no edition found")
		}
		logger.Error("Error trying get edition", err, zap.String("journey", "GetParticipant Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}

	layout := "2006-01-02"
	t, err := time.Parse(layout, closing_date)
	if err != nil {
		logger.Error("Error trying parseDate", err, zap.String("journey", "GetParticipant Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}

	t = t.AddDate(0, 0, 10)
	t = time.Date(t.Year(), t.Month(), t.Day(), 16, 0, 0, 0, t.Location())

	if !time.Now().After(t) {
		month := monthNames[int(t.Month())-1]
		dia := t.Day()
		hora := t.Format("15:04")
		msg := fmt.Sprintf("Os resultados serão liberados em %d de %s às %s", dia, month, hora)
		return &participant_response.GetParticipant{
			Particiapants:    nil,
			ClosingDate:      closing_date,
			VideoReleaseTime: msg,
			More:             false,
		}, nil
	}

	getParticipantRequest.EditionId = edition_id
	args = nil

	args = append(args, getParticipantRequest.EditionId)
	moreDataArgs := []any{getParticipantRequest.EditionId}
	from := "FROM participant AS p JOIN user_games AS u ON p.user_id = u.user_id JOIN challenge AS c ON p.edition_id = c.edition_id AND c.category = p.category AND c.sex = p.sex WHERE p.edition_id = ? AND p.checked IS true AND p.sent IS true AND desqualified IS NULL AND "
	moreData := "SELECT 1 " + from
	query = "SELECT p.video_id, u.user_id, u.name, p.category, p.noreps, p.sex, u.user, p.edition_id, p.user_time, p.placing, c.challenge, p.created_at " + from
	if getParticipantRequest.Category != "" {
		moreData += "p.category = ? AND "
		query += "p.category = ? AND "
		moreDataArgs = append(moreDataArgs, getParticipantRequest.Category)
		args = append(args, getParticipantRequest.Category)

		moreData += "p.sex = ? AND "
		query += "p.sex = ? AND "
		moreDataArgs = append(moreDataArgs, getParticipantRequest.Sex)
		args = append(args, getParticipantRequest.Sex)
	}
	if getParticipantRequest.NotVideoId != "" {
		moreData += "p.video_id != ? AND "
		query += "p.video_id != ? AND "
		moreDataArgs = append(moreDataArgs, getParticipantRequest.NotVideoId)
		args = append(args, getParticipantRequest.NotVideoId)
	}
	if getParticipantRequest.VideoId != "" {
		moreData += "p.video_id = ? AND "
		query += "p.video_id = ? AND "
		moreDataArgs = append(moreDataArgs, getParticipantRequest.VideoId)
		args = append(args, getParticipantRequest.VideoId)
	}
	if getParticipantRequest.CursorCreatedAt != "" {
		query += "p.created_at < ? AND "
		args = append(args, getParticipantRequest.CursorCreatedAt)
	}
	if getParticipantRequest.CursorUserTime != "" {
		query += "(p.user_time > ? OR p.user_time IS NULL) AND "
		args = append(args, getParticipantRequest.CursorUserTime)
	}
	query = query[:len(query)-4]

	order := "ORDER BY p.placing IS NULL, p.placing ASC, p.user_time IS NULL, p.user_time ASC, p.created_at DESC "
	query += order + "LIMIT 10 "

	rows, err := pr.mysql.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error("Error trying QueryContext", err, zap.String("journey", "GetParticipant Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}
	defer rows.Close()
	var participants []participant_response.Participant
	for rows.Next() {
		var video_id, user_id, name, category, sex, user, edition_id, challenge, created_at string
		var userTime, placing, noreps sql.NullString
		err = rows.Scan(&video_id, &user_id, &name, &category, &noreps, &sex, &user, &edition_id, &userTime, &placing, &challenge, &created_at)
		if err != nil {
			logger.Error("Error trying Scan", err, zap.String("journey", "GetParticipant Repository"))
			return nil, rest_err.NewInternalServerError("server error")
		}
		var userTimeValid any = nil
		var userPlacingValid any = nil
		var norepsValid any = nil

		if noreps.Valid {
			norepsValid = noreps.String
		}
		if placing.Valid {
			userPlacingValid = placing.String
		}
		if userTime.Valid {
			userTimeValid = userTime.String
		}
		participants = append(participants, participant_response.Participant{
			VideoId:    video_id,
			UserTime:   userTimeValid,
			Edition_id: edition_id,
			Placing:    userPlacingValid,
			Category:   category,
			Noreps:     norepsValid,
			Sex:        sex,
			Challenge:  challenge,
			User: participant_response.User{
				UserId: user_id,
				Name:   name,
				User:   user,
			},
			CreatedAt: created_at,
		})
	}

	if len(participants) == 0 {
		logger.Error("Error trying get participants", errors.New("not found"), zap.String("journey", "GetParticipant Repository"))
		return nil, rest_err.NewNotFoundError("no participants were found")
	}

	last := participants[len(participants)-1]
	if last.UserTime != nil {
		moreData += "(p.user_time > ? OR p.user_time IS NULL) AND "
		moreDataArgs = append(moreDataArgs, last.UserTime)
	} else {
		if getParticipantRequest.CursorUserTime != "" {
			moreData += "(p.user_time > ? OR p.user_time IS NULL) AND "
			moreDataArgs = append(moreDataArgs, getParticipantRequest.CursorUserTime)
		}
		moreData += "p.created_at < ? AND "
		moreDataArgs = append(moreDataArgs, last.CreatedAt)
	}

	moreData = moreData[:len(moreData)-4]
	moreData += order + "LIMIT 1"

	more := false
	err = pr.mysql.QueryRowContext(ctx, moreData, moreDataArgs...).Scan(&more)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error trying get participants", err, zap.String("journey", "GetParticipant Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}

	return &participant_response.GetParticipant{
		Particiapants:    participants,
		More:             more,
		ClosingDate:      "",
		VideoReleaseTime: "",
	}, nil
}
