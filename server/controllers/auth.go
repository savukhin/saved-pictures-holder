package controllers

import (
	"saved-pictures-holder/dto"
	"saved-pictures-holder/mappers"
	models "saved-pictures-holder/models"
	"saved-pictures-holder/utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Auth - Auth
func Auth(db *sqlx.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Successfully Authed!",
		})
	}
}

func Login(db *sqlx.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		login := dto.Login{}

		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := models.GetUserByUsername(db, login.Username)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		if user.Password != login.Password {
			c.JSON(400, gin.H{
				"message": "Invalid Password",
			})
			return
		}

		jwt, err := utils.GenerateJWT(user)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Successfully Logged In!",
			"jwt":     jwt,
		})
	}
}

// Register - Register
func Register(db *sqlx.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		register := dto.Register{}

		if err := c.ShouldBindJSON(&register); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		if register.Password != register.ConfirmPassword {
			c.JSON(400, gin.H{
				"message": "Passwords do not match",
			})
			return
		}

		user := mappers.UserRegisterToUser(&register)

		if err := user.CreateUser(db); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Successfully Registered!",
		})
	}
}
