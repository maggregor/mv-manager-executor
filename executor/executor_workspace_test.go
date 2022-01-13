package executor

import (
	"log"
	"testing"
)

// Test setCommand()

func TestSetCommandWorkspace1(t *testing.T) {
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", DatasetName: "mydatasetname", CmdType: "workspace", Queries: nil}
	executor := &WorkspaceExecutor{a1, ""}
	expected := "terraform workspace new myprojectid"
	executor.setCommand()
	if executor.Command != expected {
		log.Fatalf("Command is '%v', expected '%v'", executor.Command, expected)
	}
}

func TestSetCommandWorkspace2(t *testing.T) {
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", DatasetName: "mydatasetname", CmdType: "apply", Queries: nil}
	executor := &WorkspaceExecutor{a1, ""}
	expected := "terraform workspace select myprojectid"
	executor.setCommand()
	if executor.Command != expected {
		log.Fatalf("Command is '%v', expected '%v'", executor.Command, expected)
	}
}
