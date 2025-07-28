package judge_request

type Signin struct {
	Email string `json:"email" binding:"required,max=255"`
	Senha string `json:"senha" binding:"required,max=72"`
	Token string `json:"token" binding:"required,max=10000"`
}