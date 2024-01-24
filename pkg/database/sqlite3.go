package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func NewSQLite3Connection() (*sql.DB, error) {

	path := os.Getenv("DB_PATH")
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
