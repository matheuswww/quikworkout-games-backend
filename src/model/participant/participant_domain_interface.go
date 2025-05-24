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
	SetChecked(bool)
	GetChecked() bool
	GetCreatedAt() string
	SetCreatedAt(string)
}

func NewParticipantDomain(
	videoID, userID, editionID string,
	userTime *time.Time,
	createdAt string,
	checked bool,
) ParticipantDomainInterface {
	return &participant{
		videoID:        videoID,
		userID:         userID,
		editionID:      editionID,
		userTime: 			userTime,
		checked:         checked,
		createdAt:      createdAt,
	}
}
