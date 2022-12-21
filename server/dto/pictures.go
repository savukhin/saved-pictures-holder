package dto

import "mime/multipart"

type CreatePictureQuery struct {
	FolderID int                   `form:"folder_id" binding:"required,numeric"`
	Picture  *multipart.FileHeader `form:"picture"`
}
