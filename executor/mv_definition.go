package executor

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type MvDefinition struct {
	Id        string
	IdInt     int64
	ProjectId string
	DatasetId string
	Query     string
	FullID    string
	MVQuery   string
}

func NewMvDefinition(id string, idInt int64, projectId string, datasetId string, query string) MvDefinition {
	fullID := fmt.Sprintf("%s:%s.%s", projectId, datasetId, id)
	mvq := fmt.Sprintf("CREATE MATERIALIZED VIEW IF NOT EXISTS `%s.%s` AS %s", datasetId, id, query)
	mv := MvDefinition{id, idInt, projectId, datasetId, query, fullID, mvq}
	return mv
}

func MvDefinitionEqual(a, b MvDefinition) bool {
	return a.Id == b.Id && a.IdInt == b.IdInt && a.ProjectId == b.ProjectId &&
		a.DatasetId == b.DatasetId && a.Query == b.Query && a.FullID == b.FullID && a.MVQuery == b.MVQuery
}

func (mv MvDefinition) buildMVQuery() string {
	q := fmt.Sprintf("CREATE MATERIALIZED VIEW IF NOT EXISTS `%s.%s` AS %s", mv.DatasetId, mv.Id, mv.Query)
	return q
}

func (mv MvDefinition) isMMVExists(c *bigquery.Client, ctx context.Context) (bool, error) {
	d := c.DatasetInProject(mv.ProjectId, mv.DatasetId)
	it := d.Tables(ctx)
	for {
		t, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return false, err
		}
		md, _ := t.Metadata(ctx)
		if md.Type == "MATERIALIZED_VIEW" && md.MaterializedView != nil && md.MaterializedView.Query == mv.Query && md.FullID == mv.FullID {
			return true, nil
		}
	}
	return false, nil
}
