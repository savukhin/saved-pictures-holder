package controllers

import (
	"saved-pictures-holder/dto"
	"saved-pictures-holder/mappers"
	"saved-pictures-holder/models"
	"saved-pictures-holder/utils"
	"strconv"

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

		result := utils.ConvertToMap(folder)
		result["message"] = "Folder created"

		c.JSON(200, result)

	}
}

func GetFolderByID(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		folder_id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		folder, err := models.GetFolderByID(db, folder_id)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, _ := extractUser(c)

		if folder.UserID != user.ID {
			c.JSON(403, gin.H{
				"message": "You are not allowed to access this folder",
			})
			return
		}

		result := utils.ConvertToMap(folder)
		result["message"] = "Folder found"

		c.JSON(200, result)
	}
}
