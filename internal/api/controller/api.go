// Package controller implemnets the APIs
package controller

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
	controllerResource *ControllerResource
}

// NewAPI configures and returns application API - based on a function
// store and agent discovery service
func NewAPI(db *sql.DB) (*API, error) {
	controllerResource := NewControllerResource(database.NewFunctionStore(db))

	api := &API{
		controllerResource: controllerResource,
	}
	return api, nil
}

// Router provides application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/v1/controller", a.controllerResource.router())

	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
