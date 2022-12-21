package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Folder struct {
	ID        int          `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	UserID    int          `json:"user_id" db:"user_id"`
	CreatedAt string       `json:"created_at" db:"created_at"`
	UpdatedAt string       `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (folder *Folder) CreateFolder(db *sqlx.DB) error {
	err := db.QueryRow("INSERT INTO folders (name, user_id) VALUES ($1, $2) RETURNING id, created_at, updated_at, deleted_at",
		folder.Name, folder.UserID).
		Scan(&folder.ID, &folder.CreatedAt, &folder.UpdatedAt, &folder.DeletedAt)

	if err != nil {
		return err
	}

	return nil
}

func GetFolderByID(db *sqlx.DB, id int) (*Folder, error) {
	folder := &Folder{}

	err := db.Get(folder, "SELECT * FROM folders WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	return folder, nil
}
