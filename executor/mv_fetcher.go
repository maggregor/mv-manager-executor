package executor

import (
	"context"
	"strconv"
	"strings"

	"cloud.google.com/go/bigquery"
)

const (
	MATERIALIZED_VIEW = "MATERIALIZED_VIEW"
	TABLE             = "TABLE"
)

type FetchedTable struct {
	fullID    string
	tableName string
	projectID string
	datasetID string
	tableType string
	mvQuery   string
}

func FetchedTableEqual(a, b FetchedTable) bool {
	return a.fullID == b.fullID && a.tableName == b.tableName && a.projectID == b.projectID &&
		a.datasetID == b.datasetID && a.tableType == b.tableType && a.mvQuery == b.mvQuery
}

func fetchTables(c *bigquery.Client, ctx context.Context, projectID string, datasetID string) *bigquery.TableIterator {
	q := c.DatasetInProject(projectID, datasetID).Tables(ctx)
	return q
}

func toFetchedTable(it *bigquery.Table, ctx context.Context) *FetchedTable {
	ft := new(FetchedTable)
	md, _ := it.Metadata(ctx)
	ft.fullID = md.FullID
	ft.tableName = strings.Split(ft.fullID, ".")[1]
	ft.projectID = it.ProjectID
	ft.datasetID = it.DatasetID
	ft.tableType = string(md.Type)
	if ft.tableType == MATERIALIZED_VIEW {
		ft.mvQuery = md.MaterializedView.Query
	}
	return ft
}

func IsPositiveInt(s string) bool {
	if i, err := strconv.Atoi(s); err == nil {
		return i > 0
	}
	return false
}

func (ft *FetchedTable) isAchilioMv() bool {
	return ft.tableType == MATERIALIZED_VIEW && strings.HasPrefix(ft.tableName, "mvm_") && IsPositiveInt(strings.Split(ft.tableName, "_")[1])
}

func (ft *FetchedTable) toMvDefinition() MvDefinition {
	if !ft.isAchilioMv() {
		return MvDefinition{}
	}
	idInt, _ := strconv.ParseInt(strings.Split(ft.tableName, "_")[1], 10, 64)
	// Handle error ?

	mv := NewMvDefinition(ft.tableName, idInt, ft.projectID, ft.datasetID, ft.mvQuery)
	return mv
}
