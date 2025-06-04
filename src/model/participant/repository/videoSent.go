package participant_repository

import (
	"github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"
	model_util "github.com/matheuswww/quikworkout-games-backend/src/model/util"
)

func (pr *participantRepository) VideoSent(videoId, userId string) *rest_err.RestErr {
	return model_util.VideoSent(pr.mysql, videoId, userId)
}