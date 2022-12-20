package controllers

import (
	"saved-pictures-holder/dto"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func CreateFolder(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		folder := &dto.CreateFolderQuery{}

		if err := c.ShouldBindJSON(folder); err != nil {
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
