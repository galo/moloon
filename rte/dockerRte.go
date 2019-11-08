package rte

import (
	"github.com/galo/moloon/models"
	"github.com/galo/moloon/rte/environment"
	"log"
)

// Docker runtime implmnetations
type DockerRte struct {
}


func NewDockerRuntime() *DockerRte {
	return &DockerRte{};
}

//Executes a function in the docker inside docker runtime
func (r *DockerRte) Execute(f models.Function) (err error) {

	switch f.Spec.Lang {
	case "docker":
		return r.executeDocker(f)
	case "":
		log.Fatal("The function language description isn;t supported", f.Spec.Lang)
		return ErrFunctionSpec
	default:
		log.Fatal("The function language description isn;t supported", f.Spec.Lang)
		return ErrFunctionSpec
	}
}

// Execute docker container
func (r *DockerRte) executeDocker(f models.Function) (err error) {
	environment.CommandImplementation.Run("echo", []string{"hello works"})

	return nil
}
