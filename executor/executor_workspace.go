package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type WorkspaceExecutor struct {
	Attributes Attributes
	Command    string
}

// ExecuteShell uses the command attribute and execute it in the shell
func (e *WorkspaceExecutor) executeShell() error {
	e.setCommand()
	// TODO: Uncomment next line when terraform module is added
	// executeCommand(e.Command)
	fmt.Printf("Executing: %q\n", e.Command)
	return nil
}

func (e *WorkspaceExecutor) executeCommand() error {
	commandArray := strings.Split(e.Command, " ")
	cmd := exec.Command(commandArray[0], commandArray...)
	cmd.Dir = os.Getenv("TERRAFORM_DIR")
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", stdoutStderr)
	if err != nil {
		return err
	}
	return nil
}

func (e *WorkspaceExecutor) setQueries() {
	// WorkspaceExecutor does not use this method
	return
}

func (e *WorkspaceExecutor) setCommand() {
	e.Command = "terraform workspace"
	if e.Attributes.CmdType == WORKSPACE {
		e.Command += " new"
	} else if e.Attributes.CmdType == APPLY {
		// This should never be something else for now
		e.Command += " select"
	}
	e.Command += " " + e.Attributes.ProjectID
}
