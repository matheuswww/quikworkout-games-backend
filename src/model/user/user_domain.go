package user_domain

type user struct {
	id 			 string
	name		 string
	user     string
	dob      string
	category string
	earnings int
	cpf      string
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

func (u *user) GetDOB() string {
	return u.dob
}

func (u *user) SetDOB(dob string) {
	u.dob = dob
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