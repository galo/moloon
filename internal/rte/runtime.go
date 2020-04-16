package rte

import (
	"errors"

	"github.com/galo/moloon/pkg/models"
)

// Runtime errors
var (
	ErrFunctionSpec = errors.New("function specification error")
)

// The runtime enviroment interface
type Runtime interface {
	Execute(models.Function) (err error)
}
