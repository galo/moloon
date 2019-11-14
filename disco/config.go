package disco

import "github.com/galo/moloon/models"

// ConfigDiscovery is a static list of Agents
type ConfigDiscovery struct {
	agentsList []*models.Agent
}

func NewConfigDiscovery(url string) *ConfigDiscovery {
	a := models.NewAgent("agent-1", url)
	al := []*models.Agent{a}

	return &ConfigDiscovery{al}
}

func (c *ConfigDiscovery) GetAll() ([]*models.Agent, error) {
	return c.agentsList, nil
}
