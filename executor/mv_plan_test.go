package executor

import (
	"log"
	"testing"
)

func TestIsPlanEqual(t *testing.T) {
	mv1 := NewMvDefinition("mvm_1234", 1234, "project-id", "dataset-id", "SELECT 1")
	mv2 := NewMvDefinition("mvm_5678", 5678, "project-id", "dataset-id", "SELECT 2")
	mv3 := NewMvDefinition("mvm_9012", 9012, "project-id", "dataset-id", "SELECT 3")
	p1 := make(plan)
	p1[mv1] = CREATE
	p1[mv2] = DELETE
	p1[mv3] = CREATE
	p2 := make(plan)
	p2[mv1] = CREATE
	p2[mv2] = DELETE
	p3 := make(plan)
	p3[mv1] = CREATE
	p3[mv2] = DELETE
	p3[mv3] = CREATE
	p4 := make(plan)
	p4[mv1] = CREATE
	p4[mv2] = DELETE
	p4[mv3] = DELETE
	if PlanEqual(p1, p2) {
		log.Fatalln("ko")
	}
	if !PlanEqual(p1, p3) {
		log.Fatalln("ko")
	}
	if PlanEqual(p1, p4) {
		log.Fatalln("ko")
	}
}

func TestMvIsInMvArray(t *testing.T) {
	mv1 := NewMvDefinition("mvm_1234", 1234, "project-id", "dataset-id", "SELECT 1")
	mv2 := NewMvDefinition("mvm_5678", 5678, "project-id", "dataset-id", "SELECT 2")
	mv3 := NewMvDefinition("mvm_9012", 9012, "project-id", "dataset-id", "SELECT 3")
	mvs := []MvDefinition{mv1, mv2}
	if !MvIsInMvArray(mv1, mvs) {
		log.Fatalln("ko")
	}
	if MvIsInMvArray(mv3, mvs) {
		log.Fatalln("ko")
	}
	mvs = []MvDefinition{}
	if MvIsInMvArray(mv1, mvs) || MvIsInMvArray(mv2, mvs) || MvIsInMvArray(mv3, mvs) {
		log.Fatalln("ko")
	}
}

func TestBuildPlan1(t *testing.T) {
	id1 := "mvm_1234"
	idInt1 := int64(1234)
	projectID1 := "project-id"
	datasetID1 := "dataset-id"
	query1 := "SELECT 1"
	id2 := "mvm_5678"
	idInt2 := int64(5678)
	projectID2 := "project-id"
	datasetID2 := "dataset-id"
	query2 := "SELECT 2"
	query3 := "SELECT 3"
	mv1 := NewMvDefinition(id1, idInt1, projectID1, datasetID1, query1)
	mv11 := NewMvDefinition(id1, idInt1, projectID1, datasetID1, query1)
	mv2 := NewMvDefinition(id2, idInt2, projectID2, datasetID2, query2)
	mv21 := NewMvDefinition(id2, idInt2, projectID2, datasetID2, query2)
	newMvs1 := []MvDefinition{mv1, mv2}
	oldMvs1 := []MvDefinition{mv11, mv21}
	bp1 := BuildPlan(newMvs1, oldMvs1)
	bpe1 := make(plan)
	if !PlanEqual(bp1, bpe1) {
		log.Fatalln("Both plan should be empty")
	}

	// Test 2
	newMvs2 := []MvDefinition{mv1}
	oldMvs2 := []MvDefinition{mv11, mv21}
	bp2 := BuildPlan(newMvs2, oldMvs2)
	bpe2 := make(plan)
	bpe2[mv21] = DELETE
	if !PlanEqual(bp2, bpe2) {
		log.Fatalf("Plan should be %v\n", bpe2)
	}
	x := bpe2[mv21]
	if bp2[mv21] != x {
		log.Fatalln("Plan should be to delete mv21")
	}

	// Test 3
	mv3 := NewMvDefinition(id2, idInt2, projectID2, datasetID2, query3)
	newMvs3 := []MvDefinition{mv1, mv3}
	oldMvs3 := []MvDefinition{mv11, mv21}
	bp3 := BuildPlan(newMvs3, oldMvs3)
	bpe3 := make(plan)
	bpe3[mv3] = CREATE
	bpe3[mv21] = DELETE
	if !PlanEqual(bp3, bpe3) {
		log.Fatalf("Plan should be %v\n", bpe3)
	}
	y := bpe3[mv21]
	z := bpe3[mv3]
	if bp3[mv21] != y && bp3[mv3] != z {
		log.Fatalln("Plan should be to delete mv21")
	}
}

func TestMvsEqual(t *testing.T) {
	mv1 := NewMvDefinition("mvm_1234", 1234, "project-id", "dataset-id", "SELECT 1")
	mv2 := NewMvDefinition("mvm_5678", 5678, "project-id", "dataset-id", "SELECT 2")
	mv3 := NewMvDefinition("mvm_9012", 9012, "project-id", "dataset-id", "SELECT 3")
	mvs1 := []MvDefinition{mv1, mv2}
	mvs2 := []MvDefinition{mv1, mv2, mv3}
	mvs3 := []MvDefinition{mv1, mv3, mv2}
	mvs4 := []MvDefinition{mv1, mv3}
	mvs11 := []MvDefinition{mv1, mv2}
	mvs22 := []MvDefinition{mv1, mv2, mv3}
	mvs33 := []MvDefinition{mv1, mv3, mv2}
	mvsempty1 := []MvDefinition{}
	mvsempty2 := []MvDefinition{}
	if MvsEqual(mvs1, mvs2) {
		log.Fatalln("ko")
	}
	if MvsEqual(mvs1, mvs22) {
		log.Fatalln("ko")
	}
	if MvsEqual(mvs11, mvs2) {
		log.Fatalln("ko")
	}
	if MvsEqual(mvs11, mvs22) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvs1, mvs11) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvs11, mvs1) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvs2, mvs22) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvs1, mvs1) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvs11, mvs11) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvs2, mvs2) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvs22, mvs22) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvsempty1, mvsempty2) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvsempty2, mvsempty1) {
		log.Fatalln("ko")
	}
	if MvsEqual(mvs1, nil) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvsempty1, nil) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvs3, mvs33) {
		log.Fatalln("ko")
	}
	if !MvsEqual(mvs3, mvs2) {
		log.Fatalln("ko")
	}
	if MvsEqual(mvs1, mvs3) {
		log.Fatalln("ko")
	}
	if MvsEqual(mvs1, mvs4) {
		log.Fatalln("ko")
	}
}
