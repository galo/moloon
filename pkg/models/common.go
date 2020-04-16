package models

// APIHeader represents the object version and kind.
type APIHeader struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

type Metadata struct {
	Name   string            `json:"name,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}
