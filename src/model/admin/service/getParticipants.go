package admin_service

import (
	"net/http"

	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/vimeo"
	admin_request "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/request"
	admin_response "github.com/matheuswww/quikworkout-games-backend/src/controller/model/admin/response"
	user_service_util "github.com/matheuswww/quikworkout-games-backend/src/model/user/service/util"
	model_util "github.com/matheuswww/quikworkout-games-backend/src/model/util"
)

func (as *adminService) GetParticipants(getParticipantsRequest *admin_request.GetParticipants) ([]admin_response.Participant, *rest_err.RestErr) {
	participants, db, restErr := as.adminRepository.GetParticipants(getParticipantsRequest)
	if restErr != nil {
		return nil, restErr
	}

	for i := 0; i < len(participants); i++ {
		resp, status, err := vimeo.GetVideo(vimeo.GetVideoParams{
			VideoID: participants[i].VideoId,
			Width: getParticipantsRequest.Width,
			Autoplay: getParticipantsRequest.Autoplay,
			Muted: getParticipantsRequest.Muted,
			Background: getParticipantsRequest.Background,
		})
		if err != nil {
			return nil, rest_err.NewInternalServerError("server error")
		}
		if status == http.StatusOK && !participants[i].Sent {
			restErr := model_util.VideoSent(db, participants[i].VideoId, participants[i].User.UserId)
			if restErr != nil {
				return nil, restErr
			}
			participants[i].Sent = true
		}
		photo, restErr := user_service_util.GetUserImage(participants[i].User.User)
		if restErr != nil {
			return nil, restErr
		}
		participants[i].User.Photo = photo
		if status == http.StatusNotFound {
			continue
		}
		participants[i].Video = resp.Html
		participants[i].Title = resp.Title
		participants[i].ThumbnailUrl = resp.ThumbnailUrl
	}
	return participants, nil
}