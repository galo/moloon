package api

import (
	"context"
	"github.com/galo/moloon/internal/controller"
	"github.com/galo/moloon/internal/database"
	"github.com/galo/moloon/internal/disco"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/spf13/viper"
)

// Server provides an http.Server.
type Server struct {
	*http.Server
}

// NewServer creates and configures an APIServer serving all application routes.
// isMaster determines the API set to  provide, and if enabled initializes the Moloon Controller
func NewServer(isMaster bool) (*Server, error) {
	log.Println("configuring server...")

	if isMaster {
		// Setup the DB
		db, err := database.DBConn()
		if err != nil {
			log.Fatal("Db cannot be configured", err)
		}

		//Initialize the Moloon Controller
		ctl := controller.GetController(database.GetFunctionStore(db), disco.NewDiscoveryService())
		ctl.Start()
	}

	api, err := New(isMaster)
	if err != nil {
		return nil, err
	}

	var addr string
	port := viper.GetString("port")

	// allow port to be set as localhost:3000 in env during development to avoid "accept incoming network connection" request on restarts
	if strings.Contains(port, ":") {
		addr = port
	} else {
		addr = ":" + port
	}

	srv := http.Server{
		Addr:    addr,
		Handler: api,
	}

	return &Server{&srv}, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *Server) Start() {
	log.Println("starting server...")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)
	// teardown logic...

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
}
