package dto

type CreateFolderQuery struct {
	Name string         `json:"name" binding:"required"`
	User CompressedUser `json:"user" binding:"required"`
}
