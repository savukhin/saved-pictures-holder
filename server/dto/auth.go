package dto

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Username        string `json:"username" binding:"required,min=5,max=32,alphanum"`
	Password        string `json:"password" binding:"required,min=8,max=32,alphanum,ascii"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Email           string `json:"email" binding:"required,email"`
}
