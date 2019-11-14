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
	switch ds := viper.GetString("discovery.config"); ds {
	case "kubernetes":
		log.Println("Setting up Kubernetes discovery")
		url := viper.GetString("kubernetes_url")
		// Kubernetes in cluster
		d = NewKubernetesDiscoveryService(url)
	case "file":
		log.Println("Setting up file based discovery")
		//url := viper.GetString("agent_url")
		// Kubernetes in cluster
		a := viper.GetString("discovery.agents")
		log.Printf("Add agent %v for discovery \n", a)
		d = NewConfigDiscovery("http://agent:3000/")
	default:
		log.Fatal("Discovery not supported ")
	}

	return d
}
