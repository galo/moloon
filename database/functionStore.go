package database

import "github.com/galo/moloon/models"

// FunctionStore defines database operations for account.
type FunctionStore interface {
	Get(name string) (*models.Function, error)
	Create(models.Function) error
	Delete(models.Function) error
	GetAll() ([]*models.Function, error)
}
