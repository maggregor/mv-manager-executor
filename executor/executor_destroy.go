package executor

import (
	"fmt"
	"log"
	"os"
)

type DestroyExecutor struct {
	Attributes Attributes
	Command    string
	VarFile    string // Absolute path to the var file to use
}

func (e *DestroyExecutor) setCommand() {
	e.Command = "terraform destroy -auto-approve"
	e.Command += " -var-file " + e.VarFile
}

func (e *DestroyExecutor) createVarFile() error {
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

func (e *DestroyExecutor) toString() string {
	var r string
	r = fmt.Sprintf("project_id = %q\n", e.Attributes.ProjectID)
	r += fmt.Sprintf("access_token = %q\n", e.Attributes.AccessToken)
	r += "mmvs = []\n"
	log.Println(r)
	return r
}
