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

func GetPictures(db *sqlx.DB, folder_id int, offset int, limit int) ([]Picture, error) {
	pictures := []Picture{}

	err := db.Select(&pictures, "SELECT * FROM pictures WHERE folder_id = $1 AND deleted_at IS NULL ORDER BY id DESC LIMIT $2 OFFSET $3", folder_id, limit, offset)

	return pictures, err
}

func GetPictureByID(db *sqlx.DB, id int) (Picture, error) {
	picture := Picture{}

	err := db.Get(&picture, "SELECT * FROM pictures WHERE id = $1", id)

	return picture, err
}

func (p *Picture) UpdatePicture(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE pictures SET title = $1, description = $2, updated_at = now()  WHERE id = $3", p.Title, p.Description, p.ID)

	return err
}

func (p *Picture) DeletePicture(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE pictures SET deleted_at = now() WHERE id = $1", p.ID)

	return err
}
