package mappers

import (
	"saved-pictures-holder/dto"
	"saved-pictures-holder/models"
)

func ToPicture(folder_id int, user *models.User, path string) *models.Picture {
	return &models.Picture{
		FolderID: folder_id,
		UserID:   user.ID,
		FileName: path,
	}
}

func ToCompressedPictures(pictures []models.Picture) []dto.CompressedPicture {
	pictures_dto := []dto.CompressedPicture{}

	for _, picture := range pictures {
		pictures_dto = append(pictures_dto, dto.CompressedPicture{
			ID:       picture.ID,
			FolderID: picture.FolderID,
			Url:      picture.FileName,
		})
	}

	return pictures_dto
}

func ToPictureResponse(pictures []models.Picture, offset int, limit int) dto.GetPicturesResponse {
	return dto.GetPicturesResponse{
		Pictures: ToCompressedPictures(pictures),
		Count:    len(pictures),
		Offset:   offset,
		Limit:    limit,
		Message:  "Pictures fetched",
	}
}
