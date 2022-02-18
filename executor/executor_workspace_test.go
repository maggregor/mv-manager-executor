package executor

import (
	"log"
	"testing"
)

// Test setCommand()

func TestSetCommandWorkspace1(t *testing.T) {
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", CmdType: "workspace", Queries: nil}
	executor := &WorkspaceExecutor{a1, ""}
	expected := "terraform workspace new -no-color myprojectid"
	executor.setCommand()
	if executor.Command != expected {
		log.Fatalf("Command is '%v', expected '%v'", executor.Command, expected)
	}
}

func TestSetCommandWorkspace2(t *testing.T) {
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", CmdType: "apply", Queries: nil}
	executor := &WorkspaceExecutor{a1, ""}
	expected := "terraform workspace select -no-color myprojectid"
	executor.setCommand()
	if executor.Command != expected {
		log.Fatalf("Command is '%v', expected '%v'", executor.Command, expected)
	}
}
