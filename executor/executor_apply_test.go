package executor

import (
	"log"
	"testing"
)

func TestSetQueries1(t *testing.T) {
	i1 := []string{"SELECT 1", "SELECT 2"}
	a1 := Attributes{AccessToken: "myAccessToken", ProjectID: "myprojectid", RegionID: "myregionid", DatasetID: "mydatasetid", Queries: i1}
	t1 := Terraform{Attributes: a1, Executor: nil}
	executor := &ApplyExecutor{t1.Attributes, "", nil, ""}
	executor.setQueries()
	if executor.Queries[0].QueryContent != "SELECT 1" {
		log.Fatalf("Expected 'SELECT 1', got %v\n", executor.Queries[0].QueryContent)
	}
	if executor.Queries[1].QueryContent != "SELECT 2" {
		log.Fatalf("Expected 'SELECT 2', got %v\n", executor.Queries[1].QueryContent)
	}
}
