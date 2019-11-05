// Package app ties together application resources and handlers.
package functions

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"github.com/dhax/go-base/logging"
)

type ctxKey int

const (
	ctxAccount ctxKey = iota
	ctxProfile
)

// API provides application resources and handlers.
type API struct {
	Links *FunctionsResource
}

// NewAPI configures and returns application API.
func NewAPI() (*API, error) {
	links := NewFunctionsResource()

	api := &API{
		Links: links,
	}
	return api, nil
}

// Router provides application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	//	r.Group(func(r chi.Router) {
	//		r.Use(jwt.Authenticator)
	//		r.Mount("/profile", a.Profile.router())
	//	})

	r.Group(func(r chi.Router) {
		r.Mount("/functions", a.Links.router())
	})

	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}