package executor

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON Custom unmarshall method for the message Data.
// Message Data should be a list of 1 entry map, each representing a
// Materialized view to be created by TerraformExecutor. Key is the dataset name, Value is the query statement
func (message *Message) UnmarshalJSON(data []byte) (err error) {
	var queries []QueryParameter
	messageData := struct {
		Data       []byte     `json:"data,omitempty"`
		Attributes Attributes `json:"attributes,omitempty"`
	}{}
	err = json.Unmarshal(data, &messageData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(messageData.Data, &queries)
	if err != nil {
		return err
	}
	if err = isMessageInvalid(queries); err != nil {
		return err
	}
	qs := removeDuplicateQueryParameterInArray(queries)
	message.Attributes.Queries = qs
	message.Attributes.AccessToken = messageData.Attributes.AccessToken
	message.Attributes.CmdType = messageData.Attributes.CmdType
	message.Attributes.ProjectID = messageData.Attributes.ProjectID
	return
}

// UnmarshalJSON Custom unmarshall method for Attributes to check that the payload is valid for terraform executor
func (attribute *Attributes) UnmarshalJSON(data []byte) error {
	required := struct {
		ProjectID string `json:"projectId"`
		CmdType   string `json:"cmdType"`
	}{}
	all := struct {
		AccessToken string `json:"accessToken,omitempty"`
		ProjectID   string `json:"projectId"`
		CmdType     string `json:"cmdType"`
	}{}
	err := json.Unmarshal(data, &required)
	if err != nil {
		return err
	} else if required.CmdType == "" || (required.CmdType != WORKSPACE && required.CmdType != APPLY) {
		err = fmt.Errorf("cmdType is required and must be equal to either 'workspace' or 'apply'")
		return err
	} else if required.ProjectID == "" {
		err = fmt.Errorf("projectId is required")
		return err
	} else {
		err = json.Unmarshal(data, &all)
		if err != nil {
			return err
		}
		attribute.AccessToken = all.AccessToken
		attribute.CmdType = all.CmdType
		attribute.ProjectID = all.ProjectID
	}
	return nil
}
