package models

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"log"
)

type Agent struct {
	APIHeader
	Metadata Metadata
	Spec     AgentSpec `json:"spec,omitempty"`
}

type AgentSpec struct {
	Uri string `json:"uri,omitempty"`
}

// NewFunction is a function factory that creates the barebone function
func NewAgent(name string, uri string) *Agent {
	var a = APIHeader{APIVersion: "v1", Kind: "agent"}
	var m = Metadata{name, make(map[string]string)}
	var s = AgentSpec{Uri: uri}

	return &Agent{APIHeader: a, Metadata: m, Spec: s}
}

// JSONMarshal marshals the api object into JSON format.
func (a *Agent) JSONMarshal() (data []byte, err error) {
	data, err = json.Marshal(a)
	return
}

// YamlMarshal marshals the api object into YAML format.
func (a *Agent) YamlMarshal() (data []byte, err error) {
	data, err = yaml.Marshal(a)
	return
}

// YamlUnmarshal ummarshals the YAML into an api object.
func (a *Agent) YamlUnmarshal(data []byte) (err error) {
	if err = yaml.Unmarshal(data, &a); err != nil {
		log.Println("Failed to convert yaml file into a agent object.")
	}
	return
}
