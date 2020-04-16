package rte

import (
	"log"

	"github.com/galo/moloon/internal/rte/environment"
	"github.com/galo/moloon/pkg/models"
)

// Docker runtime implmnetations
type DockerRte struct {
}

func NewDockerRuntime() *DockerRte {
	return &DockerRte{}
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

	var args []string
	args = append(args, "run")
	args = append(args, f.Spec.Image)
	environment.CommandImplementation.Run("docker", args)

	return nil
}
