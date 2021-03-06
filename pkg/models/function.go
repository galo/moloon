// Package models contains application specific entities.
package models

import (
	"encoding/json"
	"github.com/galo/moloon/pkg/rand"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

//Function holds the information in a Function
type Function struct {
	APIHeader
	Metadata  Metadata     `json:"metadata"`
	Id        string       `json:"name"`
	Namespace string       `json:"namespace"`
	Spec      FunctionSpec `json:"spec,omitempty"`
}

type FunctionSpec struct {
	Image string
	Lang  string
}

// NewFunction is a function factory that creates the bare bones function
func NewFunction(name string, namespace string, image string, lang string) *Function {
	var a = APIHeader{APIVersion: "v1", Kind: "function"}
	var m = Metadata{name, make(map[string]string)}
	var s = FunctionSpec{Image: image, Lang: lang}
	// default is teh default ns
	if len(namespace) == 0 {
		namespace = "default"
	}
	fid := FuncName(name, namespace)
	return &Function{APIHeader: a, Metadata: m, Id: fid, Namespace: namespace, Spec: s}
}

func FuncName(name string, namespace string) string {
	var fid = name + "-" + namespace + "-" + rand.String(4)
	return fid
}

// JSONMarshal marshals the api object into JSON format.
func (f *Function) JSONMarshal() (data []byte, err error) {
	data, err = json.Marshal(f)
	return
}

// YamlMarshal marshals the api object into YAML format.
func (f *Function) YamlMarshal() (data []byte, err error) {
	data, err = yaml.Marshal(f)
	return
}

// YamlUnmarshal ummarshals the YAML into an api object.
func (f *Function) YamlUnmarshal(data []byte) (err error) {
	if err = yaml.Unmarshal(data, &f); err != nil {
		logrus.Println("Failed to convert yaml file into a function object.")
	}
	return
}
