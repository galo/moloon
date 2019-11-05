// Package models contains application specific entities.
package models

//Function holds the information in a Function
type Function struct {
	Kind       string
	ApiVersion string

	Metadata Metadata
	Spec FunctionSpec `json:"spec,omitempty"`
}

type Metadata struct {
	Name   string `json:"name,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}

type FunctionSpec struct {
	Image string
	Lang  string
}
