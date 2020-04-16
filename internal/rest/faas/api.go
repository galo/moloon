// Package app ties together application resources and handlers.

package faas

import (
	"database/sql"
	"net/http"

	"github.com/galo/moloon/internal/database"
	"github.com/galo/moloon/internal/logging"
	"github.com/galo/moloon/internal/rte"
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
	FaaSRsc *FaaSResource
}

// NewAPI configures and returns application API.
func NewAPI(db *sql.DB) (*API, error) {
	functionRsrc := NewFaaSResource(database.NewFunctionStore(db),
		rte.NewDockerRuntime())

	api := &API{
		FaaSRsc: functionRsrc,
	}
	return api, nil
}

// Router provides application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/v1/faas", a.FaaSRsc.router())

	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
