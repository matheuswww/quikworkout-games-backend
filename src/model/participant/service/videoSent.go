package participant_service

import "github.com/matheuswww/quikworkout-games-backend/src/configuration/rest_err"

func (ps *participantService) VideoSent(videoId, userId string) *rest_err.RestErr {
	return ps.participantRepository.VideoSent(videoId, userId)
}