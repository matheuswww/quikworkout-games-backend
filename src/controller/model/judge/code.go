package judge_request

type Code struct {
	Codigo string `json:"codigo" binding:"required,max=8"`
	Token  string `json:"token"  binding:"required,max=10000"`
}
