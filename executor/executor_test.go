package executor

import (
	"log"
	"testing"
)

func TestInitApply(t *testing.T) {
	terra := Terraform{Message: Message{Attributes: Attributes{CmdType: "apply"}}}
	err := terra.init()
	if err != nil {
		log.Fatalf("Init should not be in error")
	}
	switch terra.Executor.(type) {
	case *ApplyExecutor:
		log.Printf("OK")
	case *WorkspaceExecutor:
		log.Fatalf("Executor is of type Workspace, should be Apply")
	default:
		log.Fatalf("Executor is of unsupported type, should be Apply")
	}
}

func TestInitWorkspace(t *testing.T) {
	terra := Terraform{Message: Message{Attributes: Attributes{CmdType: "workspace"}}}
	err := terra.init()
	if err != nil {
		log.Fatalf("Init should not be in error")
	}
	switch terra.Executor.(type) {
	case *ApplyExecutor:
		log.Fatalf("Executor is of type Apply, should be Workspace")
	case *WorkspaceExecutor:
		log.Printf("OK")
	default:
		log.Fatalf("Executor is of unsupported type, should be Workspace")
	}
}

func TestInitDestroy(t *testing.T) {
	terra := Terraform{Message: Message{Attributes: Attributes{CmdType: "destroy"}}}
	err := terra.init()
	if err != nil {
		log.Fatalf("Init should not be in error")
	}
	switch terra.Executor.(type) {
	case *ApplyExecutor:
		log.Fatalf("Executor is of type Apply, should be Destroy")
	case *WorkspaceExecutor:
		log.Fatalf("Executor is of type Workspace, should be Destroy")
	case *DestroyExecutor:
		log.Printf("OK")
	default:
		log.Fatalf("Executor is of unsupported type, should be Destroy")
	}
}

func TestInitUnknown(t *testing.T) {
	terra := Terraform{Message: Message{Attributes: Attributes{CmdType: "foo"}}}
	err := terra.init()
	if err == nil {
		log.Fatalf("Init should be in error")
	}
	switch terra.Executor.(type) {
	case *ApplyExecutor:
		log.Fatalf("Executor is of type Apply, should be in error")
	case *WorkspaceExecutor:
		log.Fatalf("Executor is of unsupported type, should be in error")
	default:
		log.Printf("OK")
	}
}

func TestIsExecutorInvalid(t *testing.T) {
	a := ApplyExecutor{Attributes{}, "", nil, "", ""}
	q1 := QueryParameter{MmvName: "achilio_1234", DatasetName: "mydataset", Statement: ""}
	q2 := QueryParameter{MmvName: "achilio_1234", DatasetName: "", Statement: "SELECT 1"}
	q3 := QueryParameter{MmvName: "", DatasetName: "mydataset", Statement: "SELECT 1"}
	q4 := QueryParameter{MmvName: "achilio_1234", DatasetName: "mydataset", Statement: "SELECT 1"}
	l1 := []QueryParameter{q1, q4}
	l2 := []QueryParameter{q2, q4}
	l3 := []QueryParameter{q3, q4}
	l4 := []QueryParameter{q4, q4}
	a.Queries = l1
	if err := isMessageInvalid(a.Queries); err == nil {
		log.Fatalf("l1 should be in error")
	}
	a.Queries = l2
	if err := isMessageInvalid(a.Queries); err == nil {
		log.Fatalf("l2 should be in error")
	}
	a.Queries = l3
	if err := isMessageInvalid(a.Queries); err == nil {
		log.Fatalf("l3 should be in error")
	}
	a.Queries = l4
	if err := isMessageInvalid(a.Queries); err != nil {
		log.Fatalf("l4 should not be in error: got %v", err)
	}
}
