package utils

import (
	"errors"
	models "saved-pictures-holder/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

var SECRET = []byte("my-super-secret")

type Token struct {
	UserID   int
	Username string
	jwt.StandardClaims
}

func GenerateJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token{
		Username: user.Username,
		UserID:   user.ID,
	})

	return token.SignedString(SECRET)
}

func ParseJWT(tokenString string) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(*Token), nil
}

func GetUserByJWT(db *sqlx.DB, tokenString string) (*models.User, error) {
	bearer := "Bearer "
	if tokenString[:len(bearer)] != bearer {
		return nil, errors.New("invalid Token")
	}

	without_bearer := tokenString[len(bearer):]

	token, err := ParseJWT(without_bearer)

	if err != nil {
		return nil, err
	}

	user, err := models.GetUserByID(db, token.UserID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
