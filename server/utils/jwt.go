package utils

import (
	models "saved-pictures-holder/models"

	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("my-super-secret")

type Token struct {
	ID       int
	Username string
	jwt.StandardClaims
}

func GenerateJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token{
		Username: user.Username,
		ID:       user.ID,
	})

	return token.SignedString(SECRET)
}
