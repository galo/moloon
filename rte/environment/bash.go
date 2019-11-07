/*
Copyright © 2018 The GDRS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package environment

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/briandowns/spinner"
	"github.com/golang/glog"
)

const defaultFailedCode = 1

// Bash implements the Command interface for the bash shell
type Bash struct {
}

// Run creates a new subprocess run a given command with arguments.
// Returns the stdout, stderr and the exitCode.
// Source: https://stackoverflow.com/questions/10385551/get-exit-code-go
//
func (b Bash) Run(app string, args []string) (stdout string, stderr string, exitCode int) {
	var outbuf, errbuf bytes.Buffer
	cmd := createCommand(app, args)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	glog.V(1).Infof("Running command %s with the following args: %s.", app, args)
	err := cmd.Run()
	stdout = outbuf.String()
	stderr = errbuf.String()

	if err != nil {
		// try to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			// This will happen (in OSX) if `name` is not available in $PATH,
			// in this situation, exit code could not be get, and stderr will be
			// empty string very likely, so we use the default fail code, and format err
			// to string and set to stderr
			glog.V(1).Infof("Could not get exit code for failed program: %v, %v.", app, args)
			exitCode = defaultFailedCode
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}

	// Add logs for debugging purposes:
	glog.V(1).Infof("Command result: %v", exitCode)
	glog.V(2).Infof("Command output: stdout: %v, stderr: %v", stdout, stderr)

	return
}

// RunSpinner creates a new subprocess run a given command with arguments.
// Returns the stdout, stderr and the exitCode.
// While the bash command is running it keeps a spinner in terminal to give some feedback to the user.
//
func (b Bash) RunSpinner(app string, args []string) (stdout string, stderr string, exitCode int) {
	// Create terminal spinner
	s := spinner.New(spinner.CharSets[33], 200*time.Millisecond)
	s.Color("yellow")
	s.Start()
	// Run bash command
	stdout, stderr, exitCode = b.Run(app, args)
	// Stop spinner
	s.Stop()

	return
}

// RunRetry runs Bash() function repeatedly until it either succeed or timeout.
//
func (b Bash) RunRetry(app string, args []string, interval int, timeout int) (err error) {
	elapsed := time.After(time.Duration(timeout) * time.Second)
	tick := time.Tick(time.Duration(interval) * time.Second)
	exitCode := 1

	// Create terminal spinner
	s := spinner.New(spinner.CharSets[33], 200*time.Millisecond)
	s.Start()
	s.Color("yellow")

	// Keep running bash until we're timed out or get exitCode 0
	for {
		select {
		// Got a timeout! fail with a timeout error
		case <-elapsed:
			err = fmt.Errorf("Timed out running %s after %d seconds", app, timeout)
			s.Stop()
			return
		// Got a tick, run bash
		case <-tick:
			_, _, exitCode = b.Run(app, args)
			// Finish it up if the exitCode is 0 (success)
			if exitCode == 0 {
				s.Stop()
				return nil
			}
		}
	}
}

// createCommand creates an instance of exec.Cmd with the gdrsctl binaries in the PATH
//
func createCommand(app string, args []string) (cmd *exec.Cmd) {
	cmd = exec.Command(app, args...)
	cmd.Env = appendGdrsBinaryDirToPath()

	return
}

// appendGdrsBinaryDirToPath adds the gdrsctl binaries folder as the first in the path so the commands executed will use
// the controlled version of binaries
func appendGdrsBinaryDirToPath() (env []string) {
	env = os.Environ()

	for i, varValue := range env {
		keyValue := strings.Split(varValue, "=")
		if strings.EqualFold(keyValue[0], "PATH") {
			keyValue[1] = fmt.Sprintf("%s/bin%c%s", GdrsHome(), os.PathListSeparator, keyValue[1])
		}
		env[i] = strings.Join(keyValue, "=")
	}

	return
}
