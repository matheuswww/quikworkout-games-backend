package user_domain

type user struct {
	id 			 string
	name		 string
	user     string
	category string
	earnings int
	cpf      string
	session_id string
}

func (u *user) GetId() string {
	return u.id
}

func (u *user) SetId(id string) {
	u.id = id
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) SetName(name string) {
	u.name = name
}

func (u *user) GetUser() string {
	return u.user
}

func (u *user) SetUser(user string) {
	u.user = user
}

func (u *user) GetCategory() string {
	return u.category
}

func (u *user) SetCategory(category string) {
	u.category = category
}

func (u *user) GetEarnings() int {
	return u.earnings
}

func (u *user) SetEarnings(earnings int) {
	u.earnings = earnings
}

func (u *user) GetCPF() string {
	return u.cpf
}

func (u *user) SetCPF(cpf string) {
	u.cpf = cpf
}

func (u *user) GetSessionId() string {
	return u.session_id
}

func (u *user) SetSessionId(sessionId string) {
	u.session_id = sessionId
}

