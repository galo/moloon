package rte

import (
	"errors"
	"github.com/galo/moloon/models"
)

// Runtime errors
var (
	ErrFunctionSpec = errors.New("function specification error")
)

// The runtime emviroment interface
type Runtime interface {
	Execute(models.Function) (err error)
}
