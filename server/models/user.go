package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
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
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE username = $1", username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (user *User) CreateUser(db *sqlx.DB) error {
	result, err := db.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id", user.Username, user.Password, user.Email)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	user.ID = int(id)

	return nil
}
