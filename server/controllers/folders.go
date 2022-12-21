package controllers

import (
	"saved-pictures-holder/dto"
	"saved-pictures-holder/mappers"
	"saved-pictures-holder/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func CreateFolder(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		folder_query := &dto.CreateFolderQuery{}
		user_value, exists := c.Get("user")

		if err := c.ShouldBindJSON(folder_query); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		if !exists {
			c.JSON(400, gin.H{
				"message": "User not found",
			})
		}

		user := user_value.(*models.User)

		folder := mappers.DtoFolderToModelFolder(folder_query, user.ID)

		if err := folder.CreateFolder(db); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Folder created",
		})

	}
}
