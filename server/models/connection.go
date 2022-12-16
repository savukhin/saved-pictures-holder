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
	queries_folder := os.Getenv("QUERIES_FOLDER")
	tables := []string{"user", "folders", "pictures"}

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
