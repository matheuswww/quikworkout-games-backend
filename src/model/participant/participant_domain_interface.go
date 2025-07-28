package participant_domain

type ParticipantDomainInterface interface {
	GetVideoID() string
	SetVideoID(string)
	GetUserID() string
	SetUserID(string)
	GetEditionID() string
	SetEditionID(string)
	GetUserTime() string
	SetUserTime(string)
	SetChecked(bool)
	GetChecked() bool
	SetSent(bool)
	GetSent() bool
	GetCategory() string
	SetCategory(string)
	GetSex() string
	SetSex(string)
	GetCreatedAt() string
	SetCreatedAt(string)
}

func NewParticipantDomain(
	videoID, userID, editionID string,
	createdAt,
	category,
	userTime,
	finalTime,
	sex string,
	checked bool,
	sent bool,
) ParticipantDomainInterface {
	return &participant{
		videoID:        videoID,
		userID:         userID,
		editionID:      editionID,
		userTime: 			userTime,
		finalTime:      finalTime,
		checked:        checked,
		sent: 					sent,
		category:				category,
		sex: 						sex,	
		createdAt:      createdAt,
	}
}
