package edition_domain

type Top struct {
	Gain int
	Top  int
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
	GetTops() []Top
	SetTops([]Top)
	GetCreatedAt() string
	SetCreatedAt(string)
}

func NewEditionDomain(id, startDate, closingDate, rules string, tops []Top, created_at string) EditionDomainInterface {
	return &edition{
		id: 				 id,
		startDate:   startDate,
		closingDate: closingDate,
		rules:       rules,
		tops:        tops,
		created_at: created_at,
	}
}
