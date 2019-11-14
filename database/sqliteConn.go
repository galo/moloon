package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"log"
)

func DBConn() (db *sql.DB, err error) {
	database, err := sql.Open("sqlite3", "/tmp/moloon.db")
	if err != nil {
		log.Fatal("Error opening db... Reason:", err)
	}

	return database, nil
}

func checkConn(db *sql.DB) error {
	err := db.Ping()
	return err
}
