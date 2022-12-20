package models

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Connection - Connection
func Connect() (*sqlx.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func LoadTables(db *sqlx.DB) error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	queries_folder := fmt.Sprintf("%s/%s/%s", dir, "models", "sql")
	// queries_folder := "./sql"
	tables := []string{"users", "folders", "pictures"}

	for _, table := range tables {
		sql, err := os.ReadFile(fmt.Sprintf("%s/%s.sql", queries_folder, table))

		if err != nil {
			return err
		}

		_, err = db.Exec(string(sql))

		if err != nil {
			return err
		}
	}

	return nil
}
