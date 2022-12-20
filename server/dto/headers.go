package dto

type TokenHeader struct {
	Token string `header:"Authorization" binding:"required,contains=Bearer"`
}
