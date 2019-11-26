package disco

import (
	"log"

	"github.com/galo/moloon/models"
	"github.com/spf13/viper"
)

// DiscoveryService the interface for discovery services backends
type DiscoveryService interface {
	GetAll() ([]*models.Agent, error)
}

// NewDiscoveryService returns the configured discovery service
func NewDiscoveryService() DiscoveryService {
	var d DiscoveryService

	// determine the discovery service
	switch ds := viper.GetString("discovery_config"); ds {
	// Kubernetes service discovery
	case "kubernetes":
		log.Println("Setting up Kubernetes discovery")
		url := viper.GetString("url")

		// Kubernetes in cluster
		d = NewKubernetesDiscoveryService(url)

	// File configured discovery
	case "file":
		log.Println("Setting up file based discovery")

		// TODO: actually use a list of agents instead of only 1
		a := viper.GetString("discovery_agents")
		log.Printf("Add agent %v for discovery \n", a)
		d = NewConfigDiscovery(a)

	default:
		log.Fatal("Discovery not supported ")
	}

	return d
}
