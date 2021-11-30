package executor

import (
	"context"
	"log"
	"testing"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

const (
	PROJECT_ID = "achilio-test"
	DATASET_ID = "executor_integration_test"
)

func TestClientWorking(t *testing.T) {
	ctx := context.Background()
	c, err := bigquery.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatalf("Client failed: %v\n", err)
	}
	defer c.Close()
}

func TestTableFetching1(t *testing.T) {
	ctx := context.Background()
	c, err := bigquery.NewClient(ctx, PROJECT_ID)

	if err != nil {
		log.Fatalf("Client failed: %v\n", err)
	}
	defer c.Close()

	it := fetchTables(c, ctx, PROJECT_ID, DATASET_ID)
	_, err = it.Next()
	if err == iterator.Done {
		log.Fatalf("ko")
	}
}

func TestToFetchTable1(t *testing.T) {
	ctx := context.Background()
	c, err := bigquery.NewClient(ctx, PROJECT_ID)

	if err != nil {
		log.Fatalf("Client failed: %v\n", err)
	}
	defer c.Close()
	it := fetchTables(c, ctx, PROJECT_ID, DATASET_ID)

	// var table bigquery.Table
	table, err := it.Next()
	if err == iterator.Done {
		log.Fatalf("No table found")
	}
	if err != nil {
		log.Fatalf("Error during iteration: %v", err)
	}
	ft := toFetchedTable(table, ctx)
	fte := FetchedTable{"achilio-test:executor_integration_test.a_table", "a_table", "achilio-test", "executor_integration_test", "TABLE", ""}
	if !FetchedTableEqual(*ft, fte) {
		log.Fatalf("FetchedTable is:\n%v\nexpected:\n%v", *ft, fte)
	}
}

func TestToFetchTestTables(t *testing.T) {
	// This is a first draft
	ctx := context.Background()
	c, err := bigquery.NewClient(ctx, PROJECT_ID)

	if err != nil {
		log.Fatalf("Client failed: %v\n", err)
	}
	defer c.Close()
	it := fetchTables(c, ctx, PROJECT_ID, DATASET_ID)

	// 1st TABLE: Normal table
	table, err := it.Next()
	if err == iterator.Done {
		log.Fatalf("No table found")
	}
	if err != nil {
		log.Fatalf("Error during iteration: %v", err)
	}
	ft1 := toFetchedTable(table, ctx)
	fte1 := FetchedTable{"achilio-test:executor_integration_test.a_table", "a_table", "achilio-test", "executor_integration_test", "TABLE", ""}
	if !FetchedTableEqual(*ft1, fte1) {
		log.Fatalf("FetchedTable is:\n%v\nexpected:\n%v", *ft1, fte1)
	}

	// 2nd TABLE: Not Achilio MV
	table, err = it.Next()
	if err == iterator.Done {
		log.Fatalf("No table found")
	}
	if err != nil {
		log.Fatalf("Error during iteration: %v", err)
	}
	ft2 := toFetchedTable(table, ctx)
	fte2 := FetchedTable{"achilio-test:executor_integration_test.b_notachiliomv", "b_notachiliomv", "achilio-test", "executor_integration_test", "MATERIALIZED_VIEW", "SELECT FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime)) AS a_1336711789, COUNT(*) AS a_410576920 FROM `achilio-test`.`executor_integration_test`.`tlc_yellow_trips_2015_small` GROUP BY a_1336711789"}
	if !FetchedTableEqual(*ft2, fte2) {
		log.Fatalf("FetchedTable is:\n%v\nexpected:\n%v", *ft2, fte2)
	}

	// 3rd TABLE: Achilio MV
	table, err = it.Next()
	if err == iterator.Done {
		log.Fatalf("No table found")
	}
	if err != nil {
		log.Fatalf("Error during iteration: %v", err)
	}
	ft3 := toFetchedTable(table, ctx)
	fte3 := FetchedTable{"achilio-test:executor_integration_test.mvm_12345", "mvm_12345", "achilio-test", "executor_integration_test", "MATERIALIZED_VIEW", "SELECT FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime)) AS a_1336711789, COUNT(*) AS a_410576920 FROM `achilio-test`.`executor_integration_test`.`tlc_yellow_trips_2015_small` GROUP BY a_1336711789"}
	if !FetchedTableEqual(*ft3, fte3) {
		log.Fatalf("FetchedTable is:\n%v\nexpected:\n%v", *ft3, fte3)
	}

	// 4th TABLE: Not Achilio MV
	table, err = it.Next()
	if err == iterator.Done {
		log.Fatalf("No table found")
	}
	if err != nil {
		log.Fatalf("Error during iteration: %v", err)
	}
	ft4 := toFetchedTable(table, ctx)
	fte4 := FetchedTable{"achilio-test:executor_integration_test.mvm_fakeachiliomv", "mvm_fakeachiliomv", "achilio-test", "executor_integration_test", "MATERIALIZED_VIEW", "SELECT FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime)) AS a_1336711789, COUNT(*) AS a_410576920 FROM `achilio-test`.`executor_integration_test`.`tlc_yellow_trips_2015_small` GROUP BY a_1336711789"}
	if !FetchedTableEqual(*ft4, fte4) {
		log.Fatalf("FetchedTable is:\n%v\nexpected:\n%v", *ft4, fte4)
	}
}

func TestIsAchilioMv1(t *testing.T) {
	ft := &FetchedTable{"achilio-test:executor_integration_test.mvm_1234", "mvm_1234", "achilio-test", "executor_integration_test", "MATERIALIZED_VIEW", "SELECT FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime)) AS a_1336711789, COUNT(*) AS a_410576920 FROM `achilio-test`.`executor_integration_test`.`tlc_yellow_trips_2015_small` GROUP BY a_1336711789"}
	if !ft.isAchilioMv() {
		log.Fatalln("isAchilioMv1 should return true")
	}
}

func TestIsAchilioMv2(t *testing.T) {
	ft := &FetchedTable{"achilio-test:executor_integration_test.a_table", "a_table", "achilio-test", "executor_integration_test", "TABLE", ""}
	if ft.isAchilioMv() {
		log.Fatalln("isAchilioMv2 should return false")
	}
}

func TestIsAchilioMv3(t *testing.T) {
	ft := &FetchedTable{"achilio-test:executor_integration_test.b_notachiliomv", "b_notachiliomv", "achilio-test", "executor_integration_test", "MATERIALIZED_VIEW", "SELECT FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime)) AS a_1336711789, COUNT(*) AS a_410576920 FROM `achilio-test`.`executor_integration_test`.`tlc_yellow_trips_2015_small` GROUP BY a_1336711789"}
	if ft.isAchilioMv() {
		log.Fatalln("isAchilioMv4 should return false")
	}
}

func TestIsAchilioMv4(t *testing.T) {
	ft := &FetchedTable{"achilio-test:executor_integration_test.mvm_fakeachiliomv", "mvm_fakeachiliomv", "achilio-test", "executor_integration_test", "MATERIALIZED_VIEW", "SELECT FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime)) AS a_1336711789, COUNT(*) AS a_410576920 FROM `achilio-test`.`executor_integration_test`.`tlc_yellow_trips_2015_small` GROUP BY a_1336711789"}
	if ft.isAchilioMv() {
		log.Fatalln("isAchilioMv4 should return false")
	}
}

func TestIsPositiveInt(t *testing.T) {
	i1 := "12345"
	i2 := "12345_"
	i3 := "12345q"
	i4 := "qbxc"
	i5 := "-12345"
	i6 := "0000"
	if !IsPositiveInt(i1) {
		log.Fatalln("IsInt i1 should be true")
	}
	if IsPositiveInt(i2) {
		log.Fatalln("IsInt i2 should be false")
	}
	if IsPositiveInt(i3) {
		log.Fatalln("IsInt i3 should be false")
	}
	if IsPositiveInt(i4) {
		log.Fatalln("IsInt i4 should be false")
	}
	if IsPositiveInt(i5) {
		log.Fatalln("IsInt i5 should be false")
	}
	if IsPositiveInt(i6) {
		log.Fatalln("IsInt i6 should be false")
	}
}

func TestToMvDefinition1(t *testing.T) {
	ft := &FetchedTable{"achilio-test:executor_integration_test.a_table", "a_table", "achilio-test", "executor_integration_test", "TABLE", ""}
	mv := ft.toMvDefinition()
	mve := MvDefinition{}
	if !MvDefinitionEqual(mv, mve) {
		log.Fatalf("Transformed MvDefinition is:\n%v\nexpected:\n%v", mv, mve)
	}

}

func TestToMvDefinition2(t *testing.T) {
	ft := &FetchedTable{"achilio-test:executor_integration_test.mvm_12345", "mvm_12345", "achilio-test", "executor_integration_test", "MATERIALIZED_VIEW", "SELECT FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime)) AS a_1336711789, COUNT(*) AS a_410576920 FROM `achilio-test`.`executor_integration_test`.`tlc_yellow_trips_2015_small` GROUP BY a_1336711789"}
	mv := ft.toMvDefinition()
	mve := NewMvDefinition("mvm_12345", 12345, "achilio-test", "executor_integration_test", "SELECT FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime)) AS a_1336711789, COUNT(*) AS a_410576920 FROM `achilio-test`.`executor_integration_test`.`tlc_yellow_trips_2015_small` GROUP BY a_1336711789")
	if !MvDefinitionEqual(mv, mve) {
		log.Fatalf("Transformed MvDefinition is:\n%v\nexpected:\n%v", mv, mve)
	}

}

func TestToMvDefinition3(t *testing.T) {
	ft := &FetchedTable{"achilio-test:executor_integration_test.mvm_fakeachiliomv", "mvm_fakeachiliomv", "achilio-test", "executor_integration_test", "MATERIALIZED_VIEW", "SELECT FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime)) AS a_1336711789, COUNT(*) AS a_410576920 FROM `achilio-test`.`executor_integration_test`.`tlc_yellow_trips_2015_small` GROUP BY a_1336711789"}
	mv := ft.toMvDefinition()
	mve := MvDefinition{}
	if !MvDefinitionEqual(mv, mve) {
		log.Fatalf("Transformed MvDefinition is:\n%v\nexpected:\n%v", mv, mve)
	}
}
