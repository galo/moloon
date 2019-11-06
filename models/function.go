// Package models contains application specific entities.
package models

import "errors"

//Function holds the information in a Function
type Function struct {
	Kind       string
	ApiVersion string

	Metadata Metadata
	Spec     FunctionSpec `json:"spec,omitempty"`
}

type Metadata struct {
	Name   string            `json:"name,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}

type FunctionSpec struct {
	Image string
	Lang  string
}

// NewFunction is a function factory that creates the barebone function
func NewFunction(name string, image string, lang string) *Function {
	var m = Metadata{name, make(map[string]string)}
	var s = FunctionSpec{Image: image, Lang: lang}

	return &Function{Kind: "function", ApiVersion: "v1", Metadata: m, Spec: s}
}

// The list of error types returned from account resource.
var (
	ErrFunctionValidation = errors.New("function validation error")
	ErrFunctionNotfound   = errors.New("function not found")
	ErrInternalError      = errors.New("unknown internal error")
)
