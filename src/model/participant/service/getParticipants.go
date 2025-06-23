package participant_service

import (
	"errors"
	"net/http"
	"time"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/logger"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/vimeo"
	participant_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/request"
	participant_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/response"
	user_service_util "github.com/matheuswww/quikworkout-games-backend/src/model/user/service/util"
	"go.uber.org/zap"
)

func (ps *participantService) GetParticipants(getParticipantRequest *participant_request.GetParticipant) (*participant_response.GetParticipant, *rest_err.RestErr) {
	participants, restErr := ps.participantRepository.GetParticipants(getParticipantRequest)
	if restErr != nil {
		return nil, restErr
	}
	format := "2006-01-02"
	closing_date_formated, err := time.Parse(format, participants.ClosingDate)
	if err != nil {
		logger.Error("Error trying Parse date", err, zap.String("journey", "GetParticipants Repository"))
		return nil, rest_err.NewInternalServerError("server error")
	}
	closing_date_formated = closing_date_formated.Add(24*time.Hour - time.Second)
	now := time.Now()
	if now.Before(closing_date_formated) {
		logger.Error("Error trying GetParticipants", errors.New("the edition has not yet been closed"), zap.String("journey", "GetParticipants Repository"))
		participants.Particiapants = nil
		return participants, nil
	}
	var filtred []participant_response.Participant
	for _,participant := range participants.Particiapants {
		resp, status, err := vimeo.GetVideo(vimeo.GetVideoParams{
			VideoID: participant.VideoId,
			Width: getParticipantRequest.Width,
			Autoplay: getParticipantRequest.Autoplay,
			Muted: getParticipantRequest.Muted,
			Background: getParticipantRequest.Background,
		})
		if err != nil {
			return nil, rest_err.NewInternalServerError("server error")
		}
		if status == http.StatusNotFound {
			continue
		}
		photo, restErr := user_service_util.GetUserImage(participant.User.UserId)
		if restErr != nil {
			return nil, restErr
		}
		participant.User.Photo = photo
		participant.Video = resp.Html
		participant.Title = resp.Title
		participant.ThumbnailUrl = resp.ThumbnailUrl
		filtred = append(filtred, participant)
	}
	if len(participants.Particiapants) == 0 {
		return nil, rest_err.NewNotFoundError("no participants were found")
	}
	participants.Particiapants = filtred
	return participants, nil
}