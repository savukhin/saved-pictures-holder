package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID        int            `json:"id"`
	Username  string         `json:"username"`
	FirstName sql.NullString `json:"first_name" db:"first_name"`
	LastName  sql.NullString `json:"last_name" db:"last_name"`
	Password  string         `json:"password"`
	Email     string         `json:"email"`
	CreatedAt string         `json:"created_at" db:"created_at"`
	UpdatedAt string         `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime   `json:"delete_at" db:"deleted_at"`
}

func GetUserByID(db *sqlx.DB, id int) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(db *sqlx.DB, username string) (*User, error) {
	user := &User{}
	err := db.Get(user, "SELECT * FROM users WHERE username = $1 LIMIT 1", username)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) CreateUser(db *sqlx.DB) error {
	err := db.
		QueryRow("INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id",
			user.Username, user.Password, user.Email).
		Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}
