package controller

import (
	context "context"
	"github.com/galo/moloon/internal/database"
	"github.com/galo/moloon/internal/disco"
	"github.com/galo/moloon/internal/logging"
	"github.com/golang/glog"
	"log"
	"sync"
	"time"
)

// Initialize things once
var once sync.Once

// Seconds between refreshes
const FREQ = 10

// Singleton controller fo the Moloon master
var (
	masterController Controller
)

type Controller struct {
	Ctx              context.Context
	Store            database.FunctionStore
	DiscoveryService disco.DiscoveryService
}

// Gets the controller, only initializes once.
func GetController(store database.FunctionStore, discoveryService disco.DiscoveryService) Controller {
	log.Println("Initializing the Moloon Controller...")
	once.Do(func() {
		masterController = Controller{
			Ctx:              context.Background(),
			Store:            store,
			DiscoveryService: discoveryService,
		}
	})
	return masterController
}

// Starts the gorutine
func (ctl *Controller) Start() {
	glog.Infoln("Starting worker goruntine...")
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
	glog.Infoln("Syncing agents...")
	// Gets all agents
	agents, err := ctl.DiscoveryService.GetAll()
	if err != nil {
		log.Fatal("Error reported by the discovery service ", err)
		return
	}

	// Get all functions
	functions, err := ctl.Store.GetAll()
	if err != nil {
		log.Fatal("Error getting functions ", err)
		return
	}

	// Create functions on each agent
	for _, f := range functions {
		// Create the function on each agent

		for _, a := range agents {
			err = a.CreateFunction(*f)
			if err != nil {
				logging.Logger.Println("Error creating agent: ", err)
				glog.Errorf("Error creating functions %v on agent %v, err %v", f, a, err)
			}
		}
	}

}
