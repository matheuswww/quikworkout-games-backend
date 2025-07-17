package edition_domain

type Top struct {
	Gain     int
	Top      int
	Category string
}

type Challenge struct {
	Sex       string
	Challenge string
	Category  string
}

type EditionDomainInterface interface {
	GetId() string
	SetId(string)
	GetStartDate() string
	SetStartDate(string)
	GetClosingDate() string
	SetClosingDate(string)
	GetRules() string
	SetRules(string)
	GetChallenge() []Challenge
	SetChallenge([]Challenge)
	GetNumber() int
	SetNumber(int)
	GetTops() []Top
	SetTops([]Top)
	GetCreatedAt() string
	SetCreatedAt(string)
}

func NewEditionDomain(id, startDate, closingDate, rules string, tops []Top, challenge []Challenge, number int, created_at string) EditionDomainInterface {
	return &edition{
		id:          id,
		startDate:   startDate,
		closingDate: closingDate,
		rules:       rules,
		challenge:   challenge,
		number:      number,
		tops:        tops,
		created_at:  created_at,
	}
}
