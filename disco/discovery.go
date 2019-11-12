package disco

import (
	"github.com/galo/moloon/models"
	"github.com/spf13/viper"
	"log"
)

// FunctionStore defines database operations for account.
type DiscoveryService interface {
	GetAll() ([]*models.Agent, error)
}

func NewDiscoveryService() DiscoveryService {
	var d DiscoveryService

	// determine the discovery service
	switch viper.GetString("discovery") {
	case "kubernetes":
		url := viper.GetString("kubernetes_url")
		// Kubernetes in cluster
		d = NewKubernetesDiscoveryService(url)
	case "config":
		url := viper.GetString("agent_list")
		// Kubernetes in cluster
		d = NewKubernetesDiscoveryService(url)
	default:
		log.Fatal("Discovery not supported ")
	}

	return d
}
