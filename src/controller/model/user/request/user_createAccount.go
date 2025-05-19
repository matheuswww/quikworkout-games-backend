package user_request

type CreateAccount struct {
	User     string    `json:"user" binding:"required,min=1,max=30"`
	DOB      string    `json:"dob" binding:"required" validate:"date"`
	Category string    `json:"category" binding:"required,max=10"`
	CPF      string    `json:"cpf" binding:"required,len=11" validate:"cpf"`
	Token 	 string 	 `json:"token" binding:"required,max=10000"`
}