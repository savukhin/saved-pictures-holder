package mappers

import (
	"saved-pictures-holder/models"
)

func ToPicture(folder_id int, user *models.User, path string) *models.Picture {
	return &models.Picture{
		FolderID: folder_id,
		UserID:   user.ID,
		FileName: path,
	}
}
