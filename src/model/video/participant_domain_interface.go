package participant_domain

import "time"

type ParticipantDomainInterface interface {
	GetVideoID() string
	SetVideoID(string)
	GetUserID() string
	SetUserID(string)
	GetEditionID() string
	SetEditionID(string)
	GetUserTime() *time.Time
	SetUserTime(*time.Time)
	SetActive(bool)
	GetActive() bool
	GetCreatedAt() string
	SetCreatedAt(string)
}

func NewParticipantDomain(
	videoID, userID, editionID string,
	userTime *time.Time,
	createdAt string,
	active bool,
) ParticipantDomainInterface {
	return &participant{
		videoID:        videoID,
		userID:         userID,
		editionID:      editionID,
		userTime: 			userTime,
		active:         active,
		createdAt:      createdAt,
	}
}
