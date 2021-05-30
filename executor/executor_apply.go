package executor

import (
	"fmt"
	"log"
	"os"
)

type ApplyExecutor struct {
	Attributes Attributes
	Command    string
	Queries    []QueryParameter
	VarFile    string // Absolute path to the var file to use
}

// ExecuteShell uses the command attribute and execute it in the shell
func (e *ApplyExecutor) executeShell() error {
	// First, select the correct workspace
	tmpE := WorkspaceExecutor{Attributes: e.Attributes}
	tmpE.setCommand()
	log.Printf("Executing: %q\n", tmpE.Command)
	// TODO: Uncomment next line when terraform module is added
	executeCommand(tmpE.Command)
	// Create the var file
	err := e.createVarFile()
	if err != nil {
		return err
	}
	// Execute the apply command
	e.setCommand()
	log.Printf("Executing: %q\n", e.Command)
	// TODO: Uncomment next line when terraform module is added
	executeCommand(e.Command)
	return nil
}

func (e *ApplyExecutor) setQueries() {
	for _, query := range e.Attributes.Queries {
		tmpQ := QueryParameter{MvmName: "mvm_" + hash(query), QueryContent: query}
		e.Queries = append(e.Queries, tmpQ)
	}
	return
}

func (e *ApplyExecutor) createVarFile() error {
	varFile, err := os.CreateTemp(os.TempDir(), "terraview-*")
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	log.Println("Creating tmp var file: " + varFile.Name())
	varFile.WriteString(e.toString())
	e.VarFile = varFile.Name()
	return nil
}

func (e *ApplyExecutor) toString() string {
	var r string
	r = fmt.Sprintf("project_id = %q\n", e.Attributes.ProjectID)
	r += fmt.Sprintf("region_id = %q\n", e.Attributes.RegionID)
	r += fmt.Sprintf("dataset_id = %q\n", e.Attributes.DatasetID)
	r += "queries = {\n"
	for _, q := range e.Queries {
		r += fmt.Sprintf("\t%q: %q,\n", q.MvmName, q.QueryContent)
	}
	r += "\t}\n"

	return r
}

func (e *ApplyExecutor) setCommand() {
	e.Command = "terraform apply -auto-approve"
	e.Command += " -var-file " + e.VarFile
}
