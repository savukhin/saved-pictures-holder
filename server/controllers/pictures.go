package controllers

import (
	"os"
	"saved-pictures-holder/mappers"
	"saved-pictures-holder/models"
	"saved-pictures-holder/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const (
	MB = 1 << 20
)

func CreatePicture(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		picture_file, err := c.FormFile("picture")

		if err != nil {
			c.JSON(400, gin.H{
				"message q": err.Error(),
			})
			return
		}

		if picture_file.Size > 5*MB {
			c.JSON(400, gin.H{
				"message w": "File is too big",
			})
			return
		}

		folder_id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{
				"message e": err.Error(),
			})
			return
		}

		user, err := extractUser(c)

		if err != nil {
			c.JSON(403, gin.H{
				"message": "Forbidden",
			})
			return
		}

		picture_path := os.Getenv("PICTURE_PATH")
		path := picture_path + picture_file.Filename

		if err := c.SaveUploadedFile(picture_file, path); err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		picture_model := mappers.ToPicture(folder_id, user, path)

		if err := picture_model.CreatePicture(db); err != nil {
			c.JSON(400, gin.H{
				"message r": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Picture created",
		})
	}
}

func GetPictures(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		folder_id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		offset, err := strconv.Atoi(c.Query("offset"))

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		limit, err := strconv.Atoi(c.Query("limit"))

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		pictures, err := models.GetPictures(db, folder_id, offset, limit)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		pictures_dto := mappers.ToPictureResponse(pictures, offset, limit)
		result, err := utils.ConvertToMap(pictures_dto)
		// result, err := json.Marshal(pictures_dto)

		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, result)
	}
}
