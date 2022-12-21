package dto

type CreateFolderQuery struct {
	Name string `json:"name" binding:"required,min=1,max=255"`
}

type UpdateFolderQuery struct {
	Name string `json:"name" binding:"required,min=1,max=255"`
}
