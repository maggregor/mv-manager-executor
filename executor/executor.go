package executor

import (
	"errors"
	"fmt"
	"hash/fnv"
	"os"
	"os/exec"
	"strings"

	"github.com/acarl005/stripansi"
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

func (t *Terraform) Execute() error {
	err := t.init()
	if err != nil {
		return err
	}
	t.Executor.setQueries()
	t.Executor.setCommand()
	if err := t.Executor.executeShell(); err != nil {
		return err
	}
	return nil
}

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
	MvmName      string
	QueryContent string
}

func hash(s string) string {
	h := fnv.New32()
	h.Write([]byte(s))
	return fmt.Sprint(h.Sum32())
}

func executeCommand(c string) error {
	commandArray := strings.Split(c, " ")
	mainCommand := commandArray[0]
	args := commandArray[1:]
	cmd := exec.Command(mainCommand, args...)
	cmd.Dir = os.Getenv("TERRAFORM_DIR")
	stdoutStderr, err := cmd.CombinedOutput()
	cleanStdoutStderr := stripansi.Strip(string(stdoutStderr))
	fmt.Println(cleanStdoutStderr)
	if err != nil {
		return err
	}
	return nil
}
