package judge_domain

type judgeModel struct {
	id       string
	name     string
	email    string
	password string
}

func (ad *judgeModel) GetId() string {
	return ad.id
}

func (ad *judgeModel) GetName() string {
	return ad.name
}

func (ad *judgeModel) GetEmail() string {
	return ad.email
}

func (ad *judgeModel) GetPassword() string {
	return ad.password
}

func (ad *judgeModel) SetName(name string) {
	ad.name = name
}

func (ad *judgeModel) SetId(id string) {
	ad.id = id
}
