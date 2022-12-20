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

func UserToCompressedUser(user *models.User) *dto.CompressedUser {
	return &dto.CompressedUser{
		ID:       user.ID,
		Username: user.Username,
	}
}
