// Package api configures an http server for administration and application resources.
package api

import (
	"github.com/galo/moloon/internal/disco"
	"net/http"
	"time"

	"github.com/galo/moloon/internal/api/faas"
	master "github.com/galo/moloon/internal/api/master"
	"github.com/galo/moloon/internal/database"

	"github.com/galo/moloon/internal/api/functions"
	"github.com/galo/moloon/internal/logging"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// New configures application resources and routes.The
// isMaster parameter determines if we run in discovery mode
// discoService determines the discovery backend to use
func New(isMaster bool) (*chi.Mux, error) {
	logger := logging.NewLogger()

	// Setup the DB
	db, err := database.DBConn()
	if err != nil {
		logger.Fatal("Db cannot be configured", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Use(logging.NewStructuredLogger(logger))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// use CORS middleware if client is not served by this api, e.g. from other domain or CDN
	// r.Use(corsConfig().Handler)
	// Functions master
	functionAPI, err := functions.NewAPI(db)
	if err != nil {
		logger.WithField("module", "agent").Error(err)
		return nil, err
	}

	// When running in master mode, activate the master API
	if isMaster {
		// get the discovery service
		d := disco.NewDiscoveryService()

		// master master
		masterAPI, err := master.NewAPI(db, d)
		if err != nil {
			logger.WithField("module", "master").Error(err)
			return nil, err
		}

		logger.WithField("module", "master").Infoln("Starting master")

		r.Group(func(r chi.Router) {
			r.Mount("/api/v1/agents", masterAPI.Router())
			r.Mount("/api/v1/functions", functionAPI.Router())
		})
	} else {
		// FaaS runtime master
		faasAPI, err := faas.NewAPI(db)
		if err != nil {
			logger.WithField("module", "agent").Error(err)
			return nil, err
		}

		logger.WithField("module", "agent").Infoln("Starting agent")

		r.Group(func(r chi.Router) {
			r.Mount("/api", functionAPI.Router())
		})

		r.Group(func(r chi.Router) {
			r.Mount("/faas", faasAPI.Router())
		})
	}

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	return r, nil
}

func corsConfig() *cors.Cors {
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	return cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           86400, // Maximum value not ignored by any of major browsers
	})
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
