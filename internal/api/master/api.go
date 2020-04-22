// Package controller implements the APIs
package master

import (
	"database/sql"
	"github.com/galo/moloon/internal/disco"
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
	controllerResource *Resource
}

// NewAPI configures and returns application API - based on a function
// store and agent discovery service
func NewAPI(db *sql.DB, d disco.DiscoveryService) (*API, error) {
	controllerResource := NewResource(database.GetFunctionStore(db), d)
	api := &API{
		controllerResource: controllerResource,
	}
	return api, nil
}

// Router provides application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/", a.controllerResource.router())

	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
