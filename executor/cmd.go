package executor

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/acarl005/stripansi"
)

// Execute builds and executes the Terraform command
func (t *Terraform) Execute() error {
	err := t.init()
	if err != nil {
		return err
	}
	t.Executor.setCommand()
	if err := t.Executor.executeShell(); err != nil {
		return err
	}
	return nil
}

// executeShell uses the command attribute and execute it in the shell
func (e *WorkspaceExecutor) executeShell() error {
	e.setCommand()
	log.Printf("Executing: %q\n", e.Command)
	if err := executeCommand(e.Command); err != nil {
		return err
	}
	return nil
}

// executeShell uses the command attribute and execute it in the shell
func (e *ApplyExecutor) executeShell() error {
	// First, select the correct workspace
	tmpE := WorkspaceExecutor{Attributes: e.Attributes}
	tmpE.setCommand()
	log.Printf("Executing: %q\n", tmpE.Command)
	if err := executeCommand(tmpE.Command); err != nil {
		return err
	}
	// Create the var file for the terraform module
	err := e.createVarFile()
	if err != nil {
		return err
	}
	// Execute the apply command
	e.setCommand()
	log.Printf("Executing: %q\n", e.Command)
	executeCommand(e.Command)
	return nil
}

// executeShell uses the command attribute and execute it in the shell
func (e *DestroyExecutor) executeShell() error {
	// First, select the correct workspace
	tmpE := WorkspaceExecutor{Attributes: e.Attributes}
	tmpE.setCommand()
	log.Printf("Executing: %q\n", tmpE.Command)
	if err := executeCommand(tmpE.Command); err != nil {
		return err
	}
	// Create the var file for the terraform module
	err := e.createVarFile()
	if err != nil {
		return err
	}
	// Execute the destroy command
	e.setCommand()
	log.Printf("Executing: %q\n", e.Command)
	executeCommand(e.Command)
	return nil
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
