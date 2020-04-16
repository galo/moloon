package models

import "errors"

// The list of error types returned from function resource.
var (
	ErrFunctionValidation = errors.New("function validation error")
	ErrFunctionNotfound   = errors.New("function not found")
	ErrAgentValidation    = errors.New("agent validation error")
	ErrInternalError      = errors.New("unknown internal error")
)
