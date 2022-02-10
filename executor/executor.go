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
	DatasetName string
	MmvName     string
	Statement   string
}

// constructorQueryParameter takes a key/value pair and build a single QueryParameter from it
// the key is the DatasetName, the value is the QueryContent, mmv_{hashe(value)} is the MmvName
func constructorQueryParameter(m map[string]string) QueryParameter {
	var qp QueryParameter
	for k, v := range m {
		qp.DatasetName = k
		qp.Statement = unescapeQuotes(v)
		qp.MmvName = getMmvName(qp.Statement)
	}
	return qp
}

func getMmvName(statement string) string {
	return "mmv_" + hash(statement)
}

func hash(s string) string {
	h := fnv.New32()
	h.Write([]byte(s))
	return fmt.Sprint(h.Sum32())
}
