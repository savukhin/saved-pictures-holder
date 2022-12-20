package models

import (
	"database/sql"
)

type Folder struct {
	ID        int          `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	UserID    int          `json:"user_id" db:"user_id"`
	CreatedAt string       `json:"created_at" db:"created_at"`
	UpdatedAt string       `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
