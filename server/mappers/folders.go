package mappers

import (
	"saved-pictures-holder/dto"
	"saved-pictures-holder/models"
)

func DtoFolderToModelFolder(dtoFolder *dto.CreateFolderQuery, userId int) (modelFolder *models.Folder) {
	return &models.Folder{
		Name:   dtoFolder.Name,
		UserID: userId,
	}
}
