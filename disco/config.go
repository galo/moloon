package disco

import (
	"github.com/galo/moloon/logging"
	"github.com/galo/moloon/models"
	"net/url"
)

// ConfigDiscovery is a static list of Agents
type ConfigDiscovery struct {
	agentsList []*models.Agent
}

func NewConfigDiscovery(agentsConfig string) *ConfigDiscovery {
	u, err := url.Parse(agentsConfig)
	if err != nil {
		logging.Logger.Errorf("Error parsing agents config %v", agentsConfig, err)
		return &ConfigDiscovery{}
	}

	a := models.NewAgent(u.Hostname(), u.String())
	al := []*models.Agent{a}

	return &ConfigDiscovery{al}
}

func (c *ConfigDiscovery) GetAll() ([]*models.Agent, error) {
	return c.agentsList, nil
}
