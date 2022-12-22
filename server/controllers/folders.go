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

		result, _ := utils.ConvertToMap(folder)
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
			c.JSON(404, gin.H{
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

		if folder.DeletedAt.Valid {
			c.JSON(410, gin.H{
				"message": "Folder has been deleted",
			})
			return
		}

		result, _ := utils.ConvertToMap(folder)
		result["message"] = "Folder found"

		c.JSON(200, result)
	}
}

func GetFolders(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := extractUser(c)

		folders, err := models.GetFolders(db, user.ID)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		result := map[string]interface{}{
			"message": "Folders found",
			"folders": make([]map[string]interface{}, 0),
		}

		for _, folder := range folders {
			converted, _ := utils.ConvertToMap(folder)
			result["folders"] = append(result["folders"].([]map[string]interface{}), converted)
		}

		c.JSON(200, result)
	}
}

func UpdateFolder(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		folder_query := &dto.UpdateFolderQuery{}
		folder_id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := c.ShouldBindJSON(folder_query); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		folder, err := models.GetFolderByID(db, folder_id)

		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		if folder.DeletedAt.Valid {
			c.JSON(410, gin.H{
				"message": "Folder has been deleted",
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

		folder.Name = folder_query.Name

		if err := folder.UpdateFolder(db); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		result, err := utils.ConvertToMap(folder)

		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		result["message"] = "Folder updated"

		c.JSON(200, result)
	}
}

func DeleteFolder(db *sqlx.DB) gin.HandlerFunc {
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
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		if folder.DeletedAt.Valid {
			c.JSON(410, gin.H{
				"message": "Folder has been deleted",
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

		if err := folder.DeleteFolder(db); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		result, _ := utils.ConvertToMap(folder)
		result["message"] = "Folder deleted"

		c.JSON(200, result)
	}
}
