package executor

import (
	"encoding/json"
	"fmt"
	"strings"
)

// UnmarshalJSON Custom unmarshall method for the message Data.
// Message Data should be a ; delimited list of SQL queries, each representing a
// Materialized view to be created by TerraformExecutor
func (message *Message) UnmarshalJSON(data []byte) (err error) {
	// m := string(data)
	messageData := struct {
		Data       []byte     `json:"data,omitempty"`
		Attributes Attributes `json:"attributes,omitempty"`
	}{}
	err = json.Unmarshal(data, &messageData)
	if err != nil {
		return
	}
	// Removing empty queries
	queries := strings.Split(string(messageData.Data), ";")
	filtered := FilterEmpty(queries, func(query string) bool {
		return query != ""
	})
	message.Attributes.Queries = filtered
	message.Attributes.AccessToken = messageData.Attributes.AccessToken
	message.Attributes.CmdType = messageData.Attributes.CmdType
	message.Attributes.ProjectID = messageData.Attributes.ProjectID
	message.Attributes.RegionID = messageData.Attributes.RegionID
	message.Attributes.DatasetName = messageData.Attributes.DatasetName
	return
}

// FilterEmpty Returns false if the string in a string array is empty
func FilterEmpty(vs []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// UnmarshalJSON Custom unmarshall method for Attributes to check that the payload is valid for terraform executor
func (attribute *Attributes) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		ProjectID   string `json:"projectId"`
		RegionID    string `json:"regionId"`
		DatasetName string `json:"datasetName"`
		CmdType     string `json:"cmdType"`
	}{}
	all := struct {
		AccessToken string `json:"accessToken,omitempty"`
		ProjectID   string `json:"projectId"`
		RegionID    string `json:"regionId"`
		DatasetName string `json:"datasetName"`
		CmdType     string `json:"cmdType"`
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
	} else if required.DatasetName == "" && required.CmdType == APPLY {
		err = fmt.Errorf("datasetName is required when command is apply")
	} else {
		err = json.Unmarshal(data, &all)
		attribute.AccessToken = all.AccessToken
		attribute.CmdType = all.CmdType
		attribute.ProjectID = all.ProjectID
		attribute.RegionID = all.RegionID
		attribute.DatasetName = all.DatasetName
	}
	return
}
