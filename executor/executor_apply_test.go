package executor

import (
	"log"
	"testing"
)

// Test setQueries()

func TestSetQueries1(t *testing.T) {
	i1 := []string{"SELECT 1", "SELECT 2"}
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", DatasetName: "mydatasetname", Queries: i1}
	executor := &ApplyExecutor{a1, "", nil, ""}
	executor.setQueries()
	if len(executor.Queries) != 2 {
		log.Fatalf("Query len is %v, expected 2", len(executor.Queries))
	} else if executor.Queries[0].QueryContent != "SELECT 1" {
		log.Fatalf("Expected 'SELECT 1', got %v\n", executor.Queries[0].QueryContent)
	} else if executor.Queries[1].QueryContent != "SELECT 2" {
		log.Fatalf("Expected 'SELECT 2', got %v\n", executor.Queries[1].QueryContent)
	}
}

// Test setCommand()

func TestSetCommandApply1(t *testing.T) {
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", DatasetName: "mydatasetname", Queries: nil}
	executor := &ApplyExecutor{a1, "", nil, "/tmp/varfile1"}
	expected := "terraform apply -auto-approve -var-file /tmp/varfile1"
	executor.setCommand()
	if executor.Command != expected {
		log.Fatalf("Command is '%v', expected '%v'", executor.Command, expected)
	}
}

// Test createVarFile()

func TestCreateVarFile1(t *testing.T) {
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", DatasetName: "mydatasetname", Queries: nil}
	executor := &ApplyExecutor{a1, "", nil, ""}
	executor.createVarFile()
	if executor.VarFile == "" {
		log.Fatalf("Var file variable was not set")
	}
}

// Test toString()

func TestToString1(t *testing.T) {
	i1 := []string{"SELECT 1", "SELECT 2"}
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", DatasetName: "mydatasetname", Queries: i1}
	q1 := QueryParameter{MmvName: "mmv_1234", QueryContent: i1[0]}
	q2 := QueryParameter{MmvName: "mmv_5678", QueryContent: i1[1]}
	q := []QueryParameter{q1, q2}
	executor := &ApplyExecutor{a1, "", q, ""}
	r1 := executor.toString()
	expected := `project_id = "myprojectid"
dataset_name = "mydatasetname"
access_token = "myAccessToken"
queries = {
	"mmv_1234": "SELECT 1",
	"mmv_5678": "SELECT 2",
	}
`
	if r1 != expected {
		log.Fatalf("Executor string is: \n%v\nexpected: \n%v", r1, expected)
	}
}
