package mappers

import (
	"saved-pictures-holder/dto"
	"saved-pictures-holder/models"
)

func UserLoginToUser(user *dto.Login) *models.User {
	return &models.User{
		Username: user.Username,
		Password: user.Password,
	}
}

func UserRegisterToUser(user *dto.Register) *models.User {
	return &models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
}
