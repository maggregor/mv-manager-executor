package executor

import (
	"log"
	"testing"
)

func TestHashName1(t *testing.T) {
	s1 := "SELECT 1"
	r1 := hash(s1)
	expected := "2813671660"
	if r1 != expected {
		log.Fatalf("Hash is %v, expected %v", r1, expected)
	}
}

func TestHashName2(t *testing.T) {
	s2 := "SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type"
	r2 := hash(s2)
	expected := "2720750463"
	if r2 != expected {
		log.Fatalf("Hash is %v, expected %v", r2, expected)
	}
}

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

func TestInitUnknown(t *testing.T) {
	terra := Terraform{Message: Message{Attributes: Attributes{CmdType: "coucou"}}}
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
