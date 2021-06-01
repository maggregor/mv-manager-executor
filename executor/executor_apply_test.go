package executor

import (
	"log"
	"testing"
)

func TestSetQueries1(t *testing.T) {
	i1 := []string{"SELECT 1", "SELECT 2"}
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", RegionID: "myregionid", DatasetID: "mydatasetid", Queries: i1}
	m1 := Message{Attributes: a1}
	t1 := Terraform{Message: m1, Executor: nil}
	executor := &ApplyExecutor{t1.Message.Attributes, "", nil, ""}
	executor.setQueries()
	if len(executor.Queries) != 2 {
		log.Fatalf("Query len is %v, expected 2", len(executor.Queries))
	} else if executor.Queries[0].QueryContent != "SELECT 1" {
		log.Fatalf("Expected 'SELECT 1', got %v\n", executor.Queries[0].QueryContent)
	} else if executor.Queries[1].QueryContent != "SELECT 2" {
		log.Fatalf("Expected 'SELECT 2', got %v\n", executor.Queries[1].QueryContent)
	}
}
