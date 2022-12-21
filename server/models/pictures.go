package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type Picture struct {
	ID          int            `json:"id" db:"id"`
	FolderID    int            `json:"folder_id" db:"folder_id"`
	FileName    string         `json:"file_name" db:"file_name"`
	Title       sql.NullString `json:"title" db:"title"`
	Description sql.NullString `json:"description" db:"description"`
	UserID      int            `json:"user_id" db:"user_id"`

	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (p *Picture) CreatePicture(db *sqlx.DB) error {
	err := db.QueryRow("INSERT INTO pictures (folder_id, file_name, user_id) VALUES ($1, $2, $3) RETURNING id",
		p.FolderID, p.FileName, p.UserID).
		Scan(&p.ID)

	return err
}
