package edition_domain

type edition struct {
	id					string
	startDate   string
	closingDate string
	rules       string
	created_at  string
	number      int
	tops        []Top
	challenge	  []Challenge
}

func (e *edition) GetId() string {
	return e.id
}

func (e *edition) SetId(id string) {
	e.id = id
}

func (e *edition) GetStartDate() string {
	return e.startDate
}

func (e *edition) SetStartDate(s string) {
	e.startDate = s
}

func (e *edition) GetClosingDate() string {
	return e.closingDate
}

func (e *edition) SetClosingDate(c string) {
	e.closingDate = c
}

func (e *edition) GetRules() string {
	return e.rules
}

func (e *edition) SetRules(r string) {
	e.rules = r
}

func (e *edition) GetChallenge() []Challenge {
	return e.challenge
}

func (e *edition) SetChallenge(c []Challenge) {
	e.challenge = c
}

func (e *edition) GetNumber() int {
	return e.number
}

func (e *edition) SetNumber(r int) {
	e.number = r
}

func (e *edition) GetTops() []Top {
	return e.tops
}

func (e *edition) SetTops(t []Top) {
	e.tops = t
}

func (e *edition) GetCreatedAt() string {
	return e.created_at
}

func (e *edition) SetCreatedAt(t string) {
	e.created_at = t
}
