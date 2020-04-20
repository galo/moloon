package models

import (
	"bytes"
	"encoding/json"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Agent struct {
	APIHeader
	Metadata Metadata
	Spec     AgentSpec `json:"spec,omitempty"`
}

type AgentSpec struct {
	Uri string `json:"uri,omitempty"`
}

// NewAgent is a function factory that creates the barebone agents
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

// YamlUnmarshal ummarshals the YAML into an agent object.
func (a *Agent) YamlUnmarshal(data []byte) (err error) {
	if err = yaml.Unmarshal(data, &a); err != nil {
		glog.Info("Failed to convert yaml file into a agent object.")
	}
	return
}

// CreateFunction creates a function on all agent
func (a *Agent) CreateFunction(function Function) (err error) {
	jsonFunction, err := function.JSONMarshal()
	if err != nil {
		return err
	}

	var createFunctionUrl = a.Spec.Uri + "/api/v1/functions"
	response, err := http.Post(createFunctionUrl, "application/json", bytes.NewBuffer(jsonFunction))
	if err != nil {
		glog.Errorf("Error: The HTTP request failed with error %s\n", err)
		return err
	} else if response.StatusCode != 200 {
		glog.Errorf("Error: The HTTP request failed with error %v\n", response.StatusCode)
		return ErrInternalError
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		glog.Infoln(string(data))
	}
	return nil
}
