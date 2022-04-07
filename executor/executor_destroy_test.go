package executor

// import (
// 	"log"
// 	"testing"
// )

// // Test setCommand()
// func TestSetCommandDestroy1(t *testing.T) {
// 	a1 := Attributes{ServiceAccount: "myAccessToken", ProjectID: "myprojectid", Queries: nil}
// 	executor := &DestroyExecutor{a1, "", "/tmp/varfile1", ""}
// 	expected := "terraform destroy -auto-approve -no-color -var-file /tmp/varfile1"
// 	executor.setCommand()
// 	if executor.Command != expected {
// 		log.Fatalf("Command is '%v', expected '%v'", executor.Command, expected)
// 	}
// }

// // Test createVarFile()
// func TestCreateDestroyVarFile1(t *testing.T) {
// 	a1 := Attributes{ServiceAccount: "myAccessToken", ProjectID: "myprojectid", Queries: nil}
// 	executor := &DestroyExecutor{a1, "", "", ""}
// 	executor.createVarFile()
// 	if executor.VarFile == "" {
// 		log.Fatalf("Var file variable was not set")
// 	}
// }

// // Test toString()
// func TestDestroyToString1(t *testing.T) {
// 	q1 := QueryParameter{DatasetName: "mydataset1", Statement: "SELECT 1", MmvName: "mmv_1234"}
// 	q2 := QueryParameter{DatasetName: "mydataset2", Statement: "SELECT 2", MmvName: "mmv_5678"}
// 	i1 := []QueryParameter{q1, q2}
// 	a1 := Attributes{ServiceAccount: "myAccessToken", ProjectID: "myprojectid", Queries: i1}
// 	executor := &DestroyExecutor{a1, "", "", ""}
// 	r1 := executor.toString()
// 	expected := `project_id = "myprojectid"
// service_account = "myAccessToken"
// mmvs = []
// `
// 	if r1 != expected {
// 		log.Fatalf("Executor string is: \n%v\nexpected: \n%v", r1, expected)
// 	}
// }
