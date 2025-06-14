package participant_service

import (
	"net/http"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/vimeo"
	participant_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/request"
	participant_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/participant/response"
	user_service_util "github.com/matheuswww/quikworkout-games-backend/src/model/user/service/util"
)

func (ps *participantService) GetParticipants(getParticipantRequest *participant_request.GetParticipant) ([]participant_response.Participant, *rest_err.RestErr) {
	participants, restErr := ps.participantRepository.GetParticipants(getParticipantRequest)
	if restErr != nil {
		return nil, restErr
	}
	var filtred []participant_response.Participant
	for _,participant := range participants {
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
		photo, restErr := user_service_util.GetUserImage(participant.User.User)
		if restErr != nil {
			return nil, restErr
		}
		participant.User.Photo = photo
		participant.Video = resp.Html
		participant.Title = resp.Title
		participant.ThumbnailUrl = resp.ThumbnailUrl
		filtred = append(filtred, participant)
	}
	if len(participants) == 0 {
		return nil, rest_err.NewNotFoundError("no participants were found")
	}
	return filtred, nil
}