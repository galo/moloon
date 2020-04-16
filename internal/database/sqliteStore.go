package database

import (
	"database/sql"
	"log"

	"github.com/galo/moloon/pkg/models"
)

type FunctionStoreSQLite struct {
	db *sql.DB
}

func NewFunctionStore(db *sql.DB) *FunctionStoreSQLite {
	setupDb(db)
	return &FunctionStoreSQLite{
		db: db,
	}
}

// Get returns function by name, uniqueness is enforced by DB
func (s *FunctionStoreSQLite) Get(fname string) (*models.Function, error) {
	row := s.db.QueryRow("SELECT fname, image, lang FROM functions  WHERE fname= $1", fname)

	var name string
	var image string
	var lang string

	switch err := row.Scan(&name, &image, &lang); err {
	case sql.ErrNoRows:
		return nil, models.ErrFunctionNotfound
	case nil:
		return models.NewFunction(name, image, lang), nil
	default:
		// Some other unknown error reading from the db
		log.Fatal("Error fetching form Db", err)
		return nil, err
	}
}

// GetAll  returns all function
func (s *FunctionStoreSQLite) GetAll() ([]*models.Function, error) {
	rows, err := s.db.Query("SELECT fname, image, lang FROM functions")
	if err != nil {
		log.Fatal("Error fetching functions from Db", err)
		return nil, err
	}
	var fns = make([]*models.Function, 0)

	for rows.Next() {
		var name string
		var image string
		var lang string

		switch err := rows.Scan(&name, &image, &lang); err {
		case sql.ErrNoRows:
			return nil, models.ErrFunctionNotfound
		default:
			f := models.NewFunction(name, image, lang)
			fns = append(fns, f)
		}
	}

	return fns, nil
}

// Delete removes a function form the DB
func (s *FunctionStoreSQLite) Delete(f models.Function) error {
	row := s.db.QueryRow("DELETE FROM functions WHERE fname= $1", f.Metadata.Name)

	switch err := row.Scan(); err {
	case sql.ErrNoRows:
		return nil
	default:
		log.Println("Error while deleting a function", err)
		return nil
	}
}

// Creates a function in the Db
func (s *FunctionStoreSQLite) Create(f models.Function) error {
	statement, err := s.db.Prepare("INSERT INTO functions (fname, image, lang) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Error inserting function in db", err)
		return err
	}

	_, err = statement.Exec(f.Metadata.Name, f.Spec.Image, f.Spec.Lang)
	if err != nil {
		log.Println("Error inserting function in db", err)
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
