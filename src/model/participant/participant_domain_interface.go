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
	SetSent(bool)
	GetSent() bool
	GetCategory() string
	SetCategory(string)
	GetCreatedAt() string
	SetCreatedAt(string)
}

func NewParticipantDomain(
	videoID, userID, editionID string,
	userTime *time.Time,
	createdAt,
	category string,
	checked bool,
	sent bool,
) ParticipantDomainInterface {
	return &participant{
		videoID:        videoID,
		userID:         userID,
		editionID:      editionID,
		userTime: 			userTime,
		checked:        checked,
		sent: 					sent,
		category:				category,
		createdAt:      createdAt,
	}
}
