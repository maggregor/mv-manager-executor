package executor

import (
	"log"
	"testing"
)

// Test setCommand()
func TestSetCommandApply1(t *testing.T) {
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", Queries: nil}
	executor := &ApplyExecutor{a1, "", nil, "/tmp/varfile1"}
	expected := "terraform apply -auto-approve -no-color -var-file /tmp/varfile1"
	executor.setCommand()
	if executor.Command != expected {
		log.Fatalf("Command is '%v', expected '%v'", executor.Command, expected)
	}
}

// Test createVarFile()
func TestApplyCreateVarFile1(t *testing.T) {
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", Queries: nil}
	executor := &ApplyExecutor{a1, "", nil, ""}
	executor.createVarFile()
	if executor.VarFile == "" {
		log.Fatalf("Var file variable was not set")
	}
}

// Test toString()
func TestApplyToString1(t *testing.T) {
	q1 := QueryParameter{DatasetName: "mydataset1", Statement: "SELECT 1", MmvName: "mmv_1234"}
	q2 := QueryParameter{DatasetName: "mydataset2", Statement: "SELECT 2", MmvName: "mmv_5678"}
	i1 := []QueryParameter{q1, q2}
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", Queries: i1}
	executor := &ApplyExecutor{a1, "", i1, ""}
	r1 := executor.toString()
	expected := `project_id = "myprojectid"
access_token = "myAccessToken"
mmvs = [
	{
		"mmv_name": "mmv_1234",
		"dataset_name": "mydataset1",
		"statement": "SELECT 1"
	},
	{
		"mmv_name": "mmv_5678",
		"dataset_name": "mydataset2",
		"statement": "SELECT 2"
	},
]
`
	if r1 != expected {
		log.Fatalf("Executor string is: \n%v\nexpected: \n%v", r1, expected)
	}
}

// Test setMmvName()
func TestSetMmvName(t *testing.T) {
	s1 := "SELECT 1"
	s2 := ""
	s3 := "SELECT * FROM mytable GROUP BY 1"
	r1 := getMmvName(s1)
	r2 := getMmvName(s2)
	r3 := getMmvName(s3)
	if r1 != "mmv_2813671660" {
		log.Fatalf("Expected mmv_2813671660, got %v", r1)
	} else if r2 != "mmv_2166136261" {
		log.Fatalf("Expected mmv_2166136261, got %v", r2)
	} else if r3 != "mmv_3459322292" {
		log.Fatalf("Expected mmv_3459322292, got %v", r3)
	}
}
