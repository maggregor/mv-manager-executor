package executor

import (
	"context"
	"log"
	"testing"

	"cloud.google.com/go/bigquery"
)

func TestNewMvDefinition(t *testing.T) {
	id := "mvm_1234"
	idInt := int64(1234)
	projectID := "project-id"
	datasetID := "dataset-id"
	query := "SELECT 1"
	mv := NewMvDefinition(id, idInt, projectID, datasetID, query)
	mvqe := "CREATE MATERIALIZED VIEW IF NOT EXISTS `dataset-id.mvm_1234` AS SELECT 1"
	fullIDe := "project-id:dataset-id.mvm_1234"
	if mv.FullID != fullIDe {
		log.Fatalf("FullID is:\n%v\nexpected:\nproject-id:dataset-id.mvm_1234\n", mv.FullID)
	}
	if mv.MVQuery != mvqe {
		log.Fatalf("Create MV Query is %v, expected %v", mv.Query, mvqe)
	}
}

func TestBuildQuery1(t *testing.T) {
	id := "mvm_1234"
	idInt := int64(1234)
	projectID := "project-id"
	datasetID := "dataset-id"
	query := "SELECT 1"
	mv := NewMvDefinition(id, idInt, projectID, datasetID, query)
	i := mv.buildMVQuery()
	e := "CREATE MATERIALIZED VIEW IF NOT EXISTS `dataset-id.mvm_1234` AS SELECT 1"
	if i != e {
		log.Fatalf("Create MV Query is %v, expected %v", i, e)
	}
}

func TestBuildQuery2(t *testing.T) {
	id := "mvm_1234"
	idInt := int64(1234)
	projectID := "project-id"
	datasetID := "dataset-id"
	query := "SELECT t1.vendor_id as v1, t2.vendor_id as v2, SUM(t1.tip_amount) as sum_sales, AVG(t2.tip_amount) as avg_sales, FROM `reproduce_inner_join_mv_1.nyc_taxi` as t1 INNER JOIN `reproduce_inner_join_mv_2.nyc_taxi` as t2 ON t1.payment_type = t2.payment_type GROUP BY t1.vendor_id, t2.vendor_id;"
	mv := NewMvDefinition(id, idInt, projectID, datasetID, query)
	i := mv.buildMVQuery()
	e := "CREATE MATERIALIZED VIEW IF NOT EXISTS `dataset-id.mvm_1234` AS SELECT t1.vendor_id as v1, t2.vendor_id as v2, SUM(t1.tip_amount) as sum_sales, AVG(t2.tip_amount) as avg_sales, FROM `reproduce_inner_join_mv_1.nyc_taxi` as t1 INNER JOIN `reproduce_inner_join_mv_2.nyc_taxi` as t2 ON t1.payment_type = t2.payment_type GROUP BY t1.vendor_id, t2.vendor_id;"
	if i != e {
		log.Fatalf("Create MV Query is\n%q\nexpected\n%q\n", i, e)
	}
}

func TestBGClient1(t *testing.T) {
	projectID := "achilio-dev"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()
	q := client.DatasetInProject(projectID, "nyc_trips")
	q.Tables(ctx)
}

func TestIsMMVExistsNot(t *testing.T) {
	id := "mvm_1234"
	idInt := int64(1234)
	projectID := "achilio-dev"
	datasetID := "nyc_trips"
	query := "SELECT 1"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()
	mv := NewMvDefinition(id, idInt, projectID, datasetID, query)

	r, err := mv.isMMVExists(client, ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if r {
		log.Fatalf("MMV should not exist")
	}
}

func TestIsMMVExistsDatasetNotFound(t *testing.T) {
	id := "mvm_1234"
	idInt := int64(1234)
	projectID := "achilio-dev"
	datasetID := "dataset-id"
	query := "SELECT 1"

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()
	mv := NewMvDefinition(id, idInt, projectID, datasetID, query)

	e := "googleapi: Error 404: Not found: Dataset achilio-dev:dataset-id, notFound"
	r, err := mv.isMMVExists(client, ctx)
	if err == nil {
		log.Fatalf("Should be in error")
	}
	if r {
		log.Fatalf("MMV should not exist")
	}
	if err.Error() != e {
		log.Fatalf("Error is\n%q\nexpected\n%q", err.Error(), e)
	}
}

func TestIsMMVExists(t *testing.T) {
	id := "mvm_1369794050"
	idInt := int64(1369794050)
	projectID := "achilio-dev"
	datasetID := "nyc_trips"
	query := "SELECT FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime)) AS a_1336711789, COUNT(*) AS a_410576920 FROM `achilio-dev`.`nyc_trips`.`tlc_yellow_trips_2015_small` GROUP BY a_1336711789"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()
	mv := NewMvDefinition(id, idInt, projectID, datasetID, query)

	r, err := mv.isMMVExists(client, ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if !r {
		log.Fatalf("MMV should exist")
	}
}
