// Package app ties together application resources and handlers.

package functions

import (
	"database/sql"
	"net/http"

	"github.com/galo/moloon/internal/database"
	"github.com/galo/moloon/internal/logging"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type ctxKey int

const (
	ctxAccount ctxKey = iota
	ctxFunction
)

// API provides application resources and handlers.
type API struct {
	FunctionRsc *Controller
}

// NewAPI configures and returns application API.
func NewAPI(db *sql.DB) (*API, error) {
	//accountStore := database.NewAccountStore(db)

	functionRsrc := NewFunctionController(database.GetFunctionStore(db))

	api := &API{
		FunctionRsc: functionRsrc,
	}
	return api, nil
}

// Router provides application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/", a.FunctionRsc.router())

	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
