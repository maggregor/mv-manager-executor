package executor

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON Custom unmarshall method for Attributes to check that the payload is valid for terraform executor
func (attribute *Attributes) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		ProjectID string   `json:"projectId"`
		RegionID  string   `json:"regionId"`
		DatasetID string   `json:"datasetId"`
		CmdType   string   `json:"cmdType"`
		Queries   []string `json:"queries,omitempty"`
	}{}
	all := struct {
		AccessToken string   `json:"accessToken,omitempty"`
		ProjectID   string   `json:"projectId"`
		RegionID    string   `json:"regionId"`
		DatasetID   string   `json:"datasetId"`
		CmdType     string   `json:"cmdType"`
		Queries     []string `json:"queries,omitempty"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if required.CmdType == "" || (required.CmdType != WORKSPACE && required.CmdType != APPLY) {
		err = fmt.Errorf("cmdType is required and must be equal to either 'workspace' or 'apply'")
		return
	} else if required.ProjectID == "" {
		err = fmt.Errorf("projectId is required")
	} else if required.RegionID == "" && required.CmdType == APPLY {
		err = fmt.Errorf("regionId is required when command is apply")
	} else if required.DatasetID == "" && required.CmdType == APPLY {
		err = fmt.Errorf("datasetId is required when command is apply")
	} else if required.Queries == nil && required.CmdType == APPLY {
		err = fmt.Errorf("queries is required when command is apply")
	} else {
		err = json.Unmarshal(data, &all)
		attribute.AccessToken = all.AccessToken
		attribute.CmdType = all.CmdType
		attribute.ProjectID = all.ProjectID
		attribute.RegionID = all.RegionID
		attribute.DatasetID = all.DatasetID
		attribute.Queries = all.Queries
	}
	return
}
