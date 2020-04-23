package database

import (
	"database/sql"
	"github.com/galo/moloon/internal/logging"
	"log"

	"github.com/galo/moloon/pkg/models"
)

var (
	store *FunctionStoreSQLite
)

type FunctionStoreSQLite struct {
	db *sql.DB
}

// Returns the singleton Store
func GetFunctionStore(db *sql.DB) *FunctionStoreSQLite {
	if store == nil {
		setupDb(db)
		store = &FunctionStoreSQLite{
			db: db,
		}
	}
	return store
}

// Get returns function by name, uniqueness is enforced by DB
func (s *FunctionStoreSQLite) Get(fname string) (*models.Function, error) {
	row := s.db.QueryRow("SELECT fid, fname, namespace, image, lang FROM functions  WHERE fid= $1", fname)

	var name string
	var image string
	var lang string
	var id string
	var namespace string

	switch err := row.Scan(&id, &name, &namespace, &image, &lang); err {
	case sql.ErrNoRows:
		return nil, models.ErrFunctionNotfound
	case nil:
		return buildFunction(id, name, namespace, image, lang), nil
	default:
		// Some other unknown error reading from the db
		logging.Logger.Errorln("Error fetching form Db", err)
		return nil, err
	}
}

// NewFunction is a function factory that creates the bare bones function
func buildFunction(id string, name string, namespace string, image string, lang string) *models.Function {
	var a = models.APIHeader{APIVersion: "v1", Kind: "function"}
	var m = models.Metadata{name, make(map[string]string)}
	var s = models.FunctionSpec{Image: image, Lang: lang}

	return &models.Function{APIHeader: a, Metadata: m, Id: id, Namespace: namespace, Spec: s}
}

// GetAll  returns all function
func (s *FunctionStoreSQLite) GetAll() ([]*models.Function, error) {
	rows, err := s.db.Query("SELECT fid, fname, namespace, image, lang FROM functions")
	if err != nil {
		logging.Logger.Errorln("Error fetching functions from Db", err)
		return nil, err
	}
	var fns = make([]*models.Function, 0)

	for rows.Next() {
		var name string
		var image string
		var lang string
		var id string
		var namespace string

		switch err := rows.Scan(&id, &name, &namespace, &image, &lang); err {
		case sql.ErrNoRows:
			return nil, models.ErrFunctionNotfound
		default:
			f := buildFunction(id, name, namespace, image, lang)
			fns = append(fns, f)
		}
	}

	return fns, nil
}

// Delete removes a function form the DB
func (s *FunctionStoreSQLite) Delete(f models.Function) error {
	row := s.db.QueryRow("DELETE FROM functions WHERE fid= $1", f.Id)

	switch err := row.Scan(); err {
	case sql.ErrNoRows:
		return nil
	default:
		logging.Logger.Errorln("Error while deleting a function", err)
		return nil
	}
}

// Creates a function in the Db
func (s *FunctionStoreSQLite) Create(f models.Function) error {
	statement, err := s.db.Prepare("INSERT INTO functions (fid, fname, namespace, image, lang) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		logging.Logger.Errorln("Error inserting function in db", err)
		return err
	}

	_, err = statement.Exec(f.Id, f.Metadata.Name, f.Namespace, f.Spec.Image, f.Spec.Lang)
	if err != nil {
		logging.Logger.Errorln("Error inserting function in db", err)
		return err
	}

	return nil
}

func setupDb(db *sql.DB) {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS " +
		"functions (id INTEGER PRIMARY KEY AUTOINCREMENT, fid TEXT UNIQUE, fname TEXT UNIQUE, namespace TEXT, image TEXT, lang TEXT)")
	if err != nil {
		log.Fatal("Error setting up the db", err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal("Error setting up the db", err)
	}
}
