package executor

type WorkspaceExecutor struct {
	Attributes Attributes
	Command    string
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
