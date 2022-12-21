package controllers

import (
	"errors"
	models "saved-pictures-holder/models"

	"github.com/gin-gonic/gin"
)

func extractUser(c *gin.Context) (*models.User, error) {
	user_value, exists := c.Get("user")

	if !exists {
		return nil, errors.New("user not found")
	}

	user := user_value.(*models.User)

	return user, nil
}
