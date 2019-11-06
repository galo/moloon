package database

import (
	"database/sql"
	"github.com/galo/moloon/models"
	"log"
)

type FunctionStore struct {
	db *sql.DB
}

func NewFunctionStore(db *sql.DB) *FunctionStore {
	return &FunctionStore{
		db: db,
	}
}

func (s *FunctionStore) Get(name string) (*models.Function, error) {
	return nil, nil
}

func (s *FunctionStore) Delete(models.Function) error {
	return nil
}

func (s *FunctionStore) Create(f models.Function) error {
	setupDb(s.db)

	statement, err := s.db.Prepare("INSERT INTO functions (fname, image, lang) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Error inserting function", err)
		return err
	}
	_, err = statement.Exec(f.Metadata.Name, f.Spec.Image, f.Spec.Lang)

	if err != nil {
		log.Println("Error inserting function", err)
		return err
	}

	return nil
}

func setupDb(db *sql.DB) {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS " +
		"functions (id INTEGER PRIMARY KEY AUTOINCREMENT, fname TEXT , image TEXT, lang TEXT)")
	if err != nil {
		log.Fatal("Error setting up the db", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal("Error setting up the db", err)
	}
}
