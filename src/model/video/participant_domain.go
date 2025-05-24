package participant_domain

import "time"

type participant struct {
	videoID        string
	userID         string
	editionID      string
	userTime 			 *time.Time
	active				 bool
	createdAt      string
}

func (p *participant) GetVideoID() string {
	return p.videoID
}

func (p *participant) SetVideoID(v string) {
	p.videoID = v
}

func (p *participant) GetUserID() string {
	return p.userID
}

func (p *participant) SetUserID(u string) {
	p.userID = u
}

func (p *participant) GetEditionID() string {
	return p.editionID
}

func (p *participant) SetEditionID(e string) {
	p.editionID = e
}

func (p *participant) GetUserTime() *time.Time {
	return p.userTime
}

func (p *participant) SetUserTime(a *time.Time) {
	p.userTime = a
}

func (p *participant) GetActive() bool {
	return p.active
}

func (p *participant) SetActive(a bool) {
	p.active = a
}

func (p *participant) GetCreatedAt() string {
	return p.createdAt
}

func (p *participant) SetCreatedAt(c string) {
	p.createdAt = c
}
