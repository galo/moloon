package environment

import (
	"fmt"
	"strings"
)

// CommandDryRun implements the Command interface and instead of executing commands
// it prints the command in stdout
type CommandDryRun struct {
}

// Run shows command
func (c *CommandDryRun) Run(app string, args []string) (stdout string, stderr string, exitCode int) {
	fmt.Printf("Command execution: %v %v\n", app, strings.Join(args, " "))
	return
}

// RunSpinner shows command
func (c *CommandDryRun) RunSpinner(app string, args []string) (stdout string, stderr string, exitCode int) {
	fmt.Printf("Command execution: %v %v\n", app, strings.Join(args, " "))
	return
}

// RunRetry shows command
func (c *CommandDryRun) RunRetry(app string, args []string, interval int, timeout int) (err error) {
	fmt.Printf("Command execution: %v %v\n", app, strings.Join(args, " "))
	return
}
