package user_domain

type UserDomainInterface interface {
	GetId() string
	SetId(string)
	GetName() string
	SetName(string)
	GetUser() string
	SetUser(string)
	GetDOB() string
	SetDOB(string)
	GetCategory() string
	SetCategory(string)
	GetEarnings() int
	SetEarnings(int)
	GetCPF() string
	SetCPF(string)
}

func NewUserDomain(id, name, userName string, dob, category string, earnings int, cpf string) UserDomainInterface {
	return &user{
		id:       id,		
		name: 		name,
		user:     userName,
		dob:      dob,
		category: category,
		earnings: earnings,
		cpf:      cpf,
	}
}
