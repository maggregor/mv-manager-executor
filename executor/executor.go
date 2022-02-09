package executor

import (
	"errors"
	"fmt"
	"hash/fnv"
)

type Terraform struct {
	Message  Message
	Executor TerraformExecutor
}

type TerraformExecutor interface {
	setQueries()
	setCommand()
	executeShell() error
}

const (
	WORKSPACE = "workspace"
	APPLY     = "apply"
)

func (t *Terraform) init() error {
	switch t.Message.Attributes.CmdType {
	case WORKSPACE:
		t.Executor = &WorkspaceExecutor{t.Message.Attributes, ""}
	case APPLY:
		t.Executor = &ApplyExecutor{t.Message.Attributes, "", nil, ""}
	default:
		return errors.New("Unsupported command " + t.Message.Attributes.CmdType)
	}
	return nil
}

type QueryParameter struct {
	MmvName      string
	QueryContent string
}

func hash(s string) string {
	h := fnv.New32()
	h.Write([]byte(s))
	return fmt.Sprint(h.Sum32())
}
