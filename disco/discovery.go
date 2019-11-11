package disco

import "github.com/galo/moloon/models"

// FunctionStore defines database operations for account.
type DiscoveryService interface {
	GetAll() ([]*models.Agent, error)
}
