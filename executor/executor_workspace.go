package executor

import (
	"log"
)

type WorkspaceExecutor struct {
	Attributes Attributes
	Command    string
}

// ExecuteShell uses the command attribute and execute it in the shell
func (e *WorkspaceExecutor) executeShell() error {
	e.setCommand()
	log.Printf("Executing: %q\n", e.Command)
	if err := executeCommand(e.Command); err != nil {
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
