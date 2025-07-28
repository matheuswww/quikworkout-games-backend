package judge_domain

type JudgeDomainInterface interface {
	GetId() string
	GetName() string
	GetEmail() string
	GetPassword() string
	SetName(name string)
	SetId(id string)
}

func NewJudgeSendSigninDomain(email, password string) JudgeDomainInterface {
	return &judgeModel{
		email:    email,
		password: password,
	}
}

func NewJudgeCheckSigninDomain(id string) JudgeDomainInterface {
	return &judgeModel{
		id: id,
	}
}