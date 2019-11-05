// Package app ties together application resources and handlers.

package functions

import (
	"github.com/galo/moloon/logging"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ctxKey int

const (
	ctxAccount ctxKey = iota
	ctxFunction
)

// API provides application resources and handlers.
type API struct {
	FunctionRsc *FunctionResource
}

// NewAPI configures and returns application API.
func NewAPI(db *pg.DB) (*API, error) {
	//accountStore := database.NewAccountStore(db)
	functionRsrc := NewFunctionResource(nil)

	api := &API{
		FunctionRsc: functionRsrc,
	}
	return api, nil
}

// Router provides application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/functions", a.FunctionRsc.router())

	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
