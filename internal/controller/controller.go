package controller

import (
	context "context"
	"sync"
	"time"

	"github.com/galo/moloon/internal/database"
	"github.com/galo/moloon/internal/disco"
	"github.com/galo/moloon/internal/logging"
	nats "github.com/nats-io/nats.go"
)

// Initialize things once
var once sync.Once

// FREQ is seconds between refreshes
const FREQ = 10

// Singleton controller fo the Moloon master
var (
	masterController Controller
)

// Controller is the master controller thread
type Controller struct {
	Ctx              context.Context
	Store            database.FunctionStore
	DiscoveryService disco.DiscoveryService
}

// GetController gets the controller, only initializes once.
func GetController(store database.FunctionStore, discoveryService disco.DiscoveryService) Controller {
	once.Do(func() {
		masterController = Controller{
			Ctx:              context.Background(),
			Store:            store,
			DiscoveryService: discoveryService,
		}
	})
	return masterController
}

// Start starts the gorutine
func (ctl *Controller) Start() {
	logging.Logger.Infoln("Starting worker goruntine...")
	go ctl.doWork()
}

// Never ending loop
func (ctl *Controller) doWork() {
	for {
		ctl.syncAgents()
		time.Sleep(FREQ * time.Second)
	}

}

// Push functions on all Agents
func (ctl *Controller) syncAgents() {
	logging.Logger.Infoln("Syncing agents...")

	// Connect to the NATS server
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	// Get all functions
	functions, err := ctl.Store.GetAll()
	if err != nil {
		logging.Logger.Errorln("Error getting functions", err)
		return
	}

	// Create functions on each agent
	for _, f := range functions {
		// Push the function into the message queue
		c.Publish("function", f)
	}

	// Close connection
	c.Close()
}
