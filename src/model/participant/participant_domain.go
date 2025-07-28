package participant_domain

type participant struct {
	videoID   string
	userID    string
	editionID string
	userTime  string
	finalTime string
	checked   bool
	sent      bool
	category  string
	sex 			string
	createdAt string
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

func (p *participant) GetUserTime() string {
	return p.userTime
}

func (p *participant) SetUserTime(a string) {
	p.userTime = a
}

func (p *participant) GetFinalTime() string {
	return p.userTime
}

func (p *participant) SetFinalTime(a string) {
	p.userTime = a
}

func (p *participant) GetChecked() bool {
	return p.checked
}

func (p *participant) SetChecked(a bool) {
	p.checked = a
}

func (p *participant) GetSent() bool {
	return p.sent
}

func (p *participant) SetSent(s bool) {
	p.sent = s
}

func (p *participant) GetCategory() string {
	return p.category
}

func (p *participant) SetCategory(c string) {
	p.category = c
}

func (p *participant) GetSex() string {
	return p.sex
}

func (p *participant) SetSex(s string) {
	p.sex = s
}

func (p *participant) GetCreatedAt() string {
	return p.createdAt
}

func (p *participant) SetCreatedAt(c string) {
	p.createdAt = c
}
