package executor

import (
	"errors"
)

type Terraform struct {
	Message  Message
	Executor TerraformExecutor
}

type TerraformExecutor interface {
	setCommand()
	executeShell() error
}

const (
	WORKSPACE = "workspace"
	APPLY     = "apply"
	DESTROY   = "destroy"
)

func (t *Terraform) init() error {
	switch t.Message.Attributes.CmdType {
	case WORKSPACE:
		t.Executor = &WorkspaceExecutor{t.Message.Attributes, ""}
	case APPLY:
		t.Executor = &ApplyExecutor{t.Message.Attributes, "", nil, "", ""}
	case DESTROY:
		t.Executor = &DestroyExecutor{t.Message.Attributes, "", "", ""}
	default:
		return errors.New("Unsupported command " + t.Message.Attributes.CmdType)
	}
	return nil
}
