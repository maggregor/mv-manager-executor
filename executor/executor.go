package executor

import (
	"errors"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Terraform struct {
	Attributes Attributes
	Executor   TerraformExecutor
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
	t.Executor.executeShell()
	return nil
}

func (t *Terraform) init() error {
	switch t.Attributes.CmdType {
	case WORKSPACE:
		t.Executor = &WorkspaceExecutor{t.Attributes, ""}
	case APPLY:
		t.Executor = &ApplyExecutor{t.Attributes, "", nil, ""}
	default:
		return errors.New("Unsupported command " + t.Attributes.CmdType)
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
	log.Println(commandArray)
	mainCommand := commandArray[0]
	args := commandArray[1:]
	cmd := exec.Command(mainCommand, args...)
	cmd.Dir = os.Getenv("TERRAFORM_DIR")
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", stdoutStderr)
	if err != nil {
		return err
	}
	return nil
}
