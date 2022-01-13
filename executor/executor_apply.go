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
	_, err = varFile.WriteString(e.toString())
	if err != nil {
		return err
	}
	e.VarFile = varFile.Name()
	return nil
}

func (e *ApplyExecutor) toString() string {
	var r string
	r = fmt.Sprintf("project_id = %q\n", e.Attributes.ProjectID)
	r += fmt.Sprintf("dataset_name = %q\n", e.Attributes.DatasetName)
	r += fmt.Sprintf("access_token = %q\n", e.Attributes.AccessToken)
	r += "queries = {\n"
	for _, q := range e.Queries {
		r += fmt.Sprintf("\t%q: %q,\n", q.MvmName, q.QueryContent)
	}
	r += "\t}\n"
	log.Println(r)
	return r
}

func (e *ApplyExecutor) setCommand() {
	e.Command = "terraform apply -auto-approve"
	e.Command += " -var-file " + e.VarFile
}
