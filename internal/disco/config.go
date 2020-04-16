package disco

import (
	"net/url"

	"github.com/galo/moloon/internal/logging"
	"github.com/galo/moloon/pkg/models"
)

// ConfigDiscovery is a static list of Agents
type ConfigDiscovery struct {
	agentsList []*models.Agent
}

// NewConfigDiscovery creates a new Discovery service based on a configuration file
func NewConfigDiscovery(agentsConfig string) *ConfigDiscovery {
	u, err := url.Parse(agentsConfig)
	if err != nil {
		logging.Logger.Errorf("Error parsing agents config %v %v", agentsConfig, err)
		return &ConfigDiscovery{}
	}

	a := models.NewAgent(u.Hostname(), u.String())
	al := []*models.Agent{a}

	return &ConfigDiscovery{al}
}

// GetAll returns all mollon agents
func (c *ConfigDiscovery) GetAll() ([]*models.Agent, error) {
	return c.agentsList, nil
}
