package dto

import "mime/multipart"

type CreatePictureQuery struct {
	FolderID int                   `form:"folder_id" binding:"required,numeric"`
	Picture  *multipart.FileHeader `form:"picture"`
}

type CompressedPicture struct {
	ID          int    `json:"id"`
	FolderID    int    `json:"folder_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type GetPicturesResponse struct {
	Pictures []CompressedPicture `json:"pictures"`
	Count    int                 `json:"count"`
	Offset   int                 `json:"offset"`
	Limit    int                 `json:"limit"`
	Message  string              `json:"message"`
}

type PictureUpdateQuery struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}
