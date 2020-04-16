package environment

// Command represents the OS level execution interface needed by Moloon
type Command interface {
	Run(app string, args []string) (stdout string, stderr string, exitCode int)
	RunSpinner(app string, args []string) (stdout string, stderr string, exitCode int)
	RunRetry(app string, args []string, interval int, timeout int) (err error)
}

// CommandImplementation sets the default command to be used by Moloon
var CommandImplementation = Command(&Bash{})
