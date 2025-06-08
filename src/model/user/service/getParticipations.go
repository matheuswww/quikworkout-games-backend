package user_service

import (

	"net/http"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/vimeo"
	user_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/request"
	user_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/user/response"
	user_domain "github.com/matheuswww/quikworkout-games-backend/src/model/user"
	model_util "github.com/matheuswww/quikworkout-games-backend/src/model/util"
)

func (us *userService) GetParticipations(user_domain user_domain.UserDomainInterface, getParticipationsRequest *user_request.GetParticipations) (*user_response.GetParticipations, *rest_err.RestErr) {
	participants, db, restErr := us.userRepository.GetParticipations(user_domain, getParticipationsRequest)
	if restErr != nil {
		return nil, restErr
	}

	for i := 0; i < len(participants.Participations); i++ {
		resp, status, err := vimeo.GetVideo(vimeo.GetVideoParams{
			VideoID: participants.Participations[i].VideoId,
			Width: getParticipationsRequest.Width,
			Autoplay: getParticipationsRequest.Autoplay,
			Muted: getParticipationsRequest.Muted,
			Background: getParticipationsRequest.Background,
		})
		if err != nil {
			return nil, rest_err.NewInternalServerError("server error")
		}
		if status == http.StatusOK && !participants.Participations[i].Sent {
			restErr := model_util.VideoSent(db, participants.Participations[i].VideoId, user_domain.GetId())
			if restErr != nil {
				return nil, restErr
			}
			participants.Participations[i].Sent = true
		}
		if status == http.StatusNotFound {
			continue
		}

		participants.Participations[i].Video = resp.Html
		participants.Participations[i].Title = resp.Title
		participants.Participations[i].ThumbnailUrl = resp.ThumbnailUrl
	}
	return participants, nil
}