package executor

type WorkspaceExecutor struct {
	Attributes Attributes
	Command    string
}

func (e *WorkspaceExecutor) setCommand() {
	e.Command = "terraform workspace"
	if e.Attributes.CmdType == WORKSPACE {
		e.Command += " new -no-color"
	} else if e.Attributes.CmdType == APPLY || e.Attributes.CmdType == DESTROY {
		e.Command += " select -no-color"
	}
	e.Command += " " + e.Attributes.ProjectID
}
