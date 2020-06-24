package disco

import (
	"log"

	"github.com/galo/moloon/pkg/models"
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
		ns := viper.GetString("discovery_namespace")

		// Kubernetes in cluster
		d = NewKubernetesDiscoveryService(url, ns)

	// File configured discovery
	case "static":
		log.Println("Setting up static based discovery")

		// TODO: actually use a list of agents instead of only 1
		a := viper.GetString("discovery_agents")
		log.Printf("Add agent %v for discovery \n", a)
		d = NewConfigDiscovery(a)

	default:
		log.Fatal("Discovery not supported ")
	}

	return d
}
