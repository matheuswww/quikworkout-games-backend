package user_domain

type UserDomainInterface interface {
	GetId() string
	SetId(string)
	GetName() string
	SetName(string)
	GetUser() string
	SetUser(string)
	GetCategory() string
	SetCategory(string)
	GetEarnings() int
	SetEarnings(int)
	GetCPF() string
	SetCPF(string)
	GetSessionId() string
	SetSessionId(string)
}

func NewUserDomain(id, name, userName string, category string, earnings int, cpf, sessionId string) UserDomainInterface {
	return &user{
		id:       id,		
		name: 		name,
		user:     userName,
		category: category,
		earnings: earnings,
		cpf:      cpf,
		session_id: sessionId,
	}
}
