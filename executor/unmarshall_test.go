package executor

import (
	"encoding/json"
	"log"
	"testing"
)

// func TestAttributeUnmarshallErrorNoProject(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallErrorNoProject")
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "workspace"
// 			},
// 			"data": "W10=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err == nil {
// 		log.Fatalln("json.Unmarshal: should be in error: no project")
// 		return
// 	} else {
// 		log.Printf("%v\n", err)
// 	}
// }

// func TestAttributeUnmarshallInvalidData(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallInvalidData")
// 	// Data is: activating project: notre-vie
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "workspace",
// 				"projectId": "myproject"
// 			},
// 			"data": "YWN0aXZhdGluZyBwcm9qZWN0OiBub3RyZS12aWU=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 		"subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err == nil {
// 		log.Fatalln("json.Unmarshal: should be in error: data is not the correct format")
// 		return
// 	} else {
// 		log.Printf("%v\n", err)
// 	}
// }

// func TestAttributeUnmarshallInvalidCmd1(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallInvalidCmd1")
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "foo"
// 			},
// 			"data": "W10=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err == nil {
// 		log.Fatalf("json.Unmarshal: should be in error: invalid cmd, got %v", err)
// 		return
// 	} else {
// 		log.Printf("%v\n", err)
// 	}
// }

// func TestAttributeUnmarshallInvalidCmd2(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallInvalidCmd2")
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "unvalidcmd",
// 				"projectId": "myproject"
// 			},
// 			"data": "W10=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err == nil {
// 		log.Fatalf("json.Unmarshal: should be in error invalid cmd, got %v\n", err)
// 		return
// 	} else {
// 		log.Printf("%v\n", err)
// 	}
// }

// func TestUnmarshallError(t *testing.T) {
// 	log.Println("TestUnmarshallError")
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "workspace",
// 				"projectId": "myproject",
// 			},
// 			"data": "W10=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err == nil {
// 		log.Fatalf("json.Unmarshal: should in error: wrong json formatting: %v\n", err)
// 	}
// }

// func TestAttributeUnmarshallWorkspace(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallWorkspace")
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "workspace",
// 				"projectId": "myproject"
// 			},
// 			"data": "W10=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err != nil {
// 		log.Fatalf("json.Unmarshal: should not be in error: %v\n", err)
// 		return
// 	} else if r1.Message.Attributes.CmdType != "workspace" {
// 		log.Fatalf("Attribute cmdType is %v, expected 'workspace'", r1.Message.Attributes.CmdType)
// 	} else if r1.Message.Attributes.ProjectID != "myproject" {
// 		log.Fatalf("Attribute projectId is %v, expected 'myproject'", r1.Message.Attributes.CmdType)
// 	} else if r1.Message.Attributes.AccessToken != "value" {
// 		log.Fatalf("Attribute accessToken is %v, expected 'value'", r1.Message.Attributes.CmdType)
// 	} else {
// 		log.Printf("%v\n", r1.Message.Attributes)
// 	}
// }

// func TestAttributeUnmarshallWorkspaceNoToken(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallWorkspaceNoToken")
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"cmdType": "workspace",
// 				"projectId": "myproject"
// 			},
// 			"data": "W10=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err != nil {
// 		log.Fatalf("json.Unmarshal: should not be in error: %v\n", err)
// 		return
// 	} else if r1.Message.Attributes.CmdType != "workspace" {
// 		log.Fatalf("Attribute cmdType is %v, expected 'workspace'", r1.Message.Attributes.CmdType)
// 	} else if r1.Message.Attributes.ProjectID != "myproject" {
// 		log.Fatalf("Attribute projectId is %v, expected 'myproject'", r1.Message.Attributes.CmdType)
// 	} else {
// 		log.Printf("%v\n", r1.Message.Attributes)
// 	}
// }

// func TestAttributeUnmarshallEmptyQueries(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallEmptyQueries")
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "apply",
// 				"projectId": "myproject"
// 			},
// 			"data": "W10=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err != nil {
// 		log.Fatalln("json.Unmarshal: should not be in error")
// 	}
// 	if len(r1.Message.Attributes.Queries) != 0 {
// 		log.Fatalf("There are %v queries, expected 0", len(r1.Message.Attributes.Queries))
// 	}
// }

// func TestAttributeUnmarshallValidData1(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallValidData1")
// 	// data is: [{"datasetName":"mydataset","mmvName":"achilio_1234","statement":"SELECT 1"}]
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "apply",
// 				"projectId": "myproject"
// 			},
// 			"data": "W3siZGF0YXNldE5hbWUiOiJteWRhdGFzZXQiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0Iiwic3RhdGVtZW50IjoiU0VMRUNUIDEifV0=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err != nil {
// 		log.Fatalf("json.Unmarshal: should not be in error, got %v\n", err)
// 		return
// 	} else if len(r1.Message.Attributes.Queries) != 1 {
// 		log.Fatalf("Queries length is %v, expected 1.\n", len(r1.Message.Attributes.Queries))
// 	} else if r1.Message.Attributes.Queries[0].Statement != "SELECT 1" {
// 		log.Fatalf("Query is %v, expected 'SELECT 1'", r1.Message.Attributes.Queries[0])
// 	} else {
// 		log.Printf("%v\n", err)
// 	}
// }

// func TestAttributeUnmarshallValidData2(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallValidData2")
// 	// data is: [{"datasetName":"mydataset1","mmvName":"achilio_1234","statement":"SELECT 1"},{"datasetName":"mydataset2","mmvName":"achilio_12345","statement":"SELECT 2"}]
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "apply",
// 				"projectId": "myproject"
// 			},
// 			"data": "W3siZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNCIsInN0YXRlbWVudCI6IlNFTEVDVCAxIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDIiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCAyIn1d",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err != nil {
// 		log.Fatalf("json.Unmarshal: should not be in error %v\n", err)
// 		return
// 	} else if len(r1.Message.Attributes.Queries) != 2 {
// 		log.Fatalf("Query len is %v, expected 2\n", len(r1.Message.Attributes.Queries))
// 	} else if r1.Message.Attributes.Queries[0].Statement != "SELECT 1" {
// 		log.Fatalf("Query 1 is %v expected 'SELECT 1'\n", r1.Message.Attributes.Queries[0].Statement)
// 	} else if r1.Message.Attributes.Queries[1].Statement != "SELECT 2" {
// 		log.Fatalf("Query 2 is %v, expected 'SELECT 2'\n", r1.Message.Attributes.Queries[1].Statement)
// 	} else if r1.Message.Attributes.Queries[0].DatasetName != "mydataset1" {
// 		log.Fatalf("Query 1 dataset is %v expected 'mydataset1'\n", r1.Message.Attributes.Queries[0].DatasetName)
// 	} else if r1.Message.Attributes.Queries[1].DatasetName != "mydataset2" {
// 		log.Fatalf("Query 2 dataset is %v, expected 'mydataset2'\n", r1.Message.Attributes.Queries[1].DatasetName)
// 	} else if r1.Message.Attributes.Queries[0].MmvName != "achilio_1234" {
// 		log.Fatalf("Query 1 name is %v expected 'achilio_1234'\n", r1.Message.Attributes.Queries[0].MmvName)
// 	} else if r1.Message.Attributes.Queries[1].MmvName != "achilio_12345" {
// 		log.Fatalf("Query 2 name is %v expected 'achilio_12345'\n", r1.Message.Attributes.Queries[1].MmvName)
// 	} else if r1.Message.Attributes.CmdType != "apply" {
// 		log.Fatalf("cmdType is %v, expected 'apply'\n", r1.Message.Attributes.CmdType)
// 	} else if r1.Message.Attributes.ProjectID != "myproject" {
// 		log.Fatalf("Attribute projectId is %v, expected 'myproject'", r1.Message.Attributes.ProjectID)
// 	} else if r1.Message.Attributes.AccessToken != "value" {
// 		log.Fatalf("Attribute accessToken is %v, expected 'value'", r1.Message.Attributes.AccessToken)
// 	} else {
// 		log.Printf("%v\n", r1.Message.Attributes)
// 	}
// }

// func TestAttributeUnmarshallApplyInvalidDataStatement(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallApplyInvalidDataStatement")
// 	// data is: [{"datasetName":"mydataset","mmvName":"achilio_1234","statement":"SELECT 1"},{"datasetName":"mydataset","mmvName":"achilio_1234"}]
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "apply",
// 				"projectId": "myproject"
// 			},
// 			"data": "W3siZGF0YXNldE5hbWUiOiJteWRhdGFzZXQiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0Iiwic3RhdGVtZW50IjoiU0VMRUNUIDEifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNCJ9XQ==",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	e := fmt.Errorf("statement is required in message for all queries")
// 	var r1 PubSubMessage
// 	err := json.Unmarshal(i2, &r1)
// 	if err == nil {
// 		log.Fatalf("json.Unmarshal: should be in error\n")
// 		return
// 	}
// 	if err.Error() != e.Error() {
// 		log.Fatalf("Expected error %v, got %v", e, err)
// 	}
// }

// func TestAttributeUnmarshallApplyInvalidDataDataset(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallApplyInvalidDataStatement")
// 	// data is: [{"datasetName":"mydataset","mmvName":"achilio_1234","statement":"SELECT 1"},{"mmvName":"achilio_1234","statement":"SELECT 2"}]
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "apply",
// 				"projectId": "myproject"
// 			},
// 			"data": "W3siZGF0YXNldE5hbWUiOiJteWRhdGFzZXQiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0Iiwic3RhdGVtZW50IjoiU0VMRUNUIDEifSx7Im1tdk5hbWUiOiJhY2hpbGlvXzEyMzQiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgMiJ9XQ==",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	e := fmt.Errorf("datasetName is required in message for all queries")
// 	var r1 PubSubMessage
// 	err := json.Unmarshal(i2, &r1)
// 	if err == nil {
// 		log.Fatalf("json.Unmarshal: should be in error\n")
// 		return
// 	}
// 	if err.Error() != e.Error() {
// 		log.Fatalf("Expected error %v, got %v", e, err)
// 	}
// }

// func TestAttributeUnmarshallApplyInvalidDataMmvName(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallApplyInvalidDataMmvName")
// 	// data is: [{"datasetName":"mydataset","mmvName":"achilio_1234","statement":"SELECT 1"},{"statement":"SELECT 2","datasetName":"mydataset"}]
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "apply",
// 				"projectId": "myproject"
// 			},
// 			"data": "W3siZGF0YXNldE5hbWUiOiJteWRhdGFzZXQiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0Iiwic3RhdGVtZW50IjoiU0VMRUNUIDEifSx7InN0YXRlbWVudCI6IlNFTEVDVCAyIiwiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQifV0=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	e := fmt.Errorf("mmvName is required in message for all queries")
// 	var r1 PubSubMessage
// 	err := json.Unmarshal(i2, &r1)
// 	if err == nil {
// 		log.Fatalf("json.Unmarshal: should be in error\n")
// 		return
// 	}
// 	if err.Error() != e.Error() {
// 		log.Fatalf("Expected error %v, got %v", e, err)
// 	}
// }

// func TestAttributeUnmarshallMessageDataEmpty(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallMessageDataEmpty")
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "apply",
// 				"projectId": "myproject"
// 			},
// 			"data": "",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err == nil {
// 		log.Fatalf("json.Unmarshal: should be in error")
// 		return
// 	} else {
// 		log.Printf("%v\n", err)
// 	}
// }

// func TestAttributeUnmarshallMessageDataEmptyArray(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallMessageDataEmptyArray")
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "apply",
// 				"projectId": "myproject"
// 			},
// 			"data": "W10=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err != nil {
// 		log.Fatalf("json.Unmarshal: should not be in error %v\n", err)
// 		return
// 	} else if len(r1.Message.Attributes.Queries) != 0 {
// 		log.Fatalf("Queries length is %v, expected 0. Empty query should be removed\n", len(r1.Message.Attributes.Queries))
// 	} else {
// 		log.Printf("%v\n", err)
// 	}
// }

// func TestAttributeUnmarshallMessageDataVeryLong(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallMessageDataVeryLong")
// 	// data is: [{"datasetName":"mydataset1","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset2","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset3","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset4","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset5","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset6","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset7","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset8","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset9","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset10","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset11","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset12","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset13","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset14","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset15","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset16","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset17","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset18","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset19","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset20","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset21","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset22","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset23","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset24","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset25","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset26","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset27","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset28","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset29","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset30","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset31","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset32","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset33","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset34","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset35","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset36","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset37","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset38","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset39","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset40","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"}]
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "apply",
// 				"projectId": "myproject"
// 			},
// 			"data": "W3siZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MiIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDMiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQ0IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0NSIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDYiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQ3IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0OCIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDkiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxMCIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDExIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MTIiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxMyIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDE0IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MTUiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxNiIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDE3IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MTgiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxOSIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDIwIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MjEiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQyMiIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDIzIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MjQiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQyNSIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDI2IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MjciLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQyOCIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDI5IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MzAiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQzMSIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDMyIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MzMiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQzNCIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDM1IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MzYiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQzNyIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDM4IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MzkiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQ0MCIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn1d",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err != nil {
// 		log.Fatalf("json.Unmarshal: should not be in error %v\n", err)
// 		return
// 	} else if len(r1.Message.Attributes.Queries) != 40 {
// 		log.Fatalf("Queries length is %v, expected 40.\n", len(r1.Message.Attributes.Queries))
// 	} else {
// 		log.Printf("%v\n", err)
// 	}
// }

// func TestAttributeUnmarshallEscapedQuotesInData(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallEscapedQuotesInData")
// 	// data is: [{"mmvName":"achilio_1234","datasetName":"mydataset","statement":"SELECT (FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime))) IN (\"December\", \"January\", \"June\", \"March\") AS a_386307744"}]
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "apply",
// 				"projectId": "myproject"
// 			},
// 			"data": "W3sibW12TmFtZSI6ImFjaGlsaW9fMTIzNCIsImRhdGFzZXROYW1lIjoibXlkYXRhc2V0Iiwic3RhdGVtZW50IjoiU0VMRUNUIChGT1JNQVRfREFURVRJTUUoXCIlQlwiLCBEQVRFVElNRShwaWNrdXBfZGF0ZXRpbWUpKSkgSU4gKFwiRGVjZW1iZXJcIiwgXCJKYW51YXJ5XCIsIFwiSnVuZVwiLCBcIk1hcmNoXCIpIEFTIGFfMzg2MzA3NzQ0In1d",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 		"subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err != nil {
// 		log.Fatalf("json.Unmarshal: should not be in error %v\n", err)
// 	}
// 	log.Printf("Query is %v", r1.Message.Attributes.Queries[0])
// 	if r1.Message.Attributes.Queries[0].Statement != `SELECT (FORMAT_DATETIME("%B", DATETIME(pickup_datetime))) IN ("December", "January", "June", "March") AS a_386307744` {
// 		log.Fatalf("Query is %v, expected %v", r1.Message.Attributes.Queries[0], `SELECT (FORMAT_DATETIME("%B", DATETIME(pickup_datetime))) IN ("December", "January", "June", "March") AS a_386307744`)
// 	}
// }

// func TestAttributeUnmarshallDuplicateQuery(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallDuplicateQuery")
// 	// data is: [{"mmvName":"achilio_1234","datasetName":"mydataset1","statement":"SELECT1"},{"mmvName":"achilio_1234","datasetName":"mydataset1","statement":"SELECT2"},{"mmvName":"achilio_1234","datasetName":"mydataset2","statement":"SELECT1"},{"mmvName":"achilio_1234","datasetName":"mydataset1","statement":"SELECT1"}]
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "apply",
// 				"projectId": "myproject"
// 			},
// 			"data": "W3sibW12TmFtZSI6ImFjaGlsaW9fMTIzNCIsImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MSIsInN0YXRlbWVudCI6IlNFTEVDVDEifSx7Im1tdk5hbWUiOiJhY2hpbGlvXzEyMzQiLCJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDEiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QyIn0seyJtbXZOYW1lIjoiYWNoaWxpb18xMjM0IiwiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQyIiwic3RhdGVtZW50IjoiU0VMRUNUMSJ9LHsibW12TmFtZSI6ImFjaGlsaW9fMTIzNCIsImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MSIsInN0YXRlbWVudCI6IlNFTEVDVDEifV0=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err != nil {
// 		log.Fatalf("json.Unmarshal: should not be in error %v\n", err)
// 	}
// 	if len(r1.Message.Attributes.Queries) != 3 {
// 		log.Fatalf("There are %v queries, expected 3", len(r1.Message.Attributes.Queries))
// 	}
// }

// func TestAttributeUnmarshallDestroy(t *testing.T) {
// 	log.Println("TestAttributeUnmarshallDestroy")
// 	i1 := `{
// 		"message": {
// 			"attributes": {
// 				"accessToken": "value",
// 				"cmdType": "destroy",
// 				"projectId": "myproject"
// 			},
// 			"data": "W10=",
// 			"messageId": "2070443601311540",
// 			"message_id": "2070443601311540",
// 			"publishTime": "2021-02-26T19:13:55.749Z",
// 			"publish_time": "2021-02-26T19:13:55.749Z"
// 		},
// 	    "subscription": "projects/myproject/subscriptions/mysubscription"
// 	}`
// 	i2 := []byte(i1)
// 	var r1 PubSubMessage
// 	if err := json.Unmarshal(i2, &r1); err != nil {
// 		log.Fatalf("json.Unmarshal: should not be in error: %v\n", err)
// 		return
// 	} else if r1.Message.Attributes.CmdType != "destroy" {
// 		log.Fatalf("Attribute cmdType is %v, expected 'destroy'", r1.Message.Attributes.CmdType)
// 	} else if r1.Message.Attributes.ProjectID != "myproject" {
// 		log.Fatalf("Attribute projectId is %v, expected 'myproject'", r1.Message.Attributes.ProjectID)
// 	} else if r1.Message.Attributes.AccessToken != "value" {
// 		log.Fatalf("Attribute accessToken is %v, expected 'value'", r1.Message.Attributes.AccessToken)
// 	} else {
// 		log.Printf("%v\n", r1.Message.Attributes)
// 	}
// }

func TestTmpMarshall(t *testing.T) {
	// {"serviceAccount":"{"type":"achilio-dev@achilio.com"}","queries":[{"mmvName":"achilio_1234","datasetName":"mydataset1","statement":"SELECT1"}]}
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "destroy",
				"projectId": "myproject"
			},
			"data": "eyJzZXJ2aWNlQWNjb3VudCI6IntcbiBcInR5cGVcIjogXCJzZXJ2aWNlX2FjY291bnRcIixcbiBc\nInByb2plY3RfaWRcIjogXCJhY2hpbGlvLWRldlwiLFxuIFwicHJpdmF0ZV9rZXlfaWRcIjogXCI5\nMTljZjNlZGUxZDAyYjM4MmNjZDM0YTliZTY4ODBjOWRkYjYxODA2XCIsXG4gXCJwcml2YXRlX2tl\neVwiOiBcIi0tLS0tQkVHSU4gUFJJVkFURSBLRVktLS0tLVxcbk1JSUV2UUlCQURBTkJna3Foa2lH\nOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRRDF2YkRPczRvc2tVODVcXG5GN21JNFNNM1Bt\nSzFrb09jeEg2cUxMdi80YVVhQXd2aWxHamRZWExZOGJkV0hDZGs4eGh0TVRoVmRTdzRNQVV0XFxu\nNjVWVFI4VWd0NjEwak5HVnBvOS8xbENxbmhrRzJwM2prUFBFNXhTdTNTZmRDcXlPMFRpNG5YS1FG\nK21zRFFWSVxcbnJ1NTJ6N3RrclNFVk55dTc0bERqSzQwSXFvNDRjdnBzVXVqeHJBZndkM1NYS0tq\ncDcvQmhyZHlTcUhWaDVWMGNcXG42NG1sdEZWVUMvTE94bUNqc1JCN0dSVHE1UG1JdFp4TjdPN2cv\nVnVjS1cwVjlGVlErTVZ4bTU4Ymp0aXlzU1gxXFxuVlJTczA0dWpEM0t0aGVadWx3Yi9GbGtGbS92\nTks0ZndtVElmdUVZbnFicUsweW8wNGZUMDI5VWpvaWpMdkd5Y1xcbmpXOEVteE83QWdNQkFBRUNn\nZ0VBUnhuVldGa1dONExybXlkRUpWaU9xZU5CQ0N3V0xIdWw4bkx6Q2p3dFd2TmFcXG5zdm9oeFR2\nMjBOSlBLM2tGNU8xNU9jNnQ3L3Z4Q1dqTGR0ZmM0clczVVpvVGM4OXlySXYrcFF4TFZySmJwMklG\nXFxub054Rnk4alljU2ZINnVxMXRyVE9ab3dYbzUxbm5NSS9yeUgrR1I1TnAvbDN3djJ0aDdVWmxr\nRDVrM0pWcC9yN1xcbldWNzY4VVg2a3lSbFQyaytLMXo5QTlxNGVwdHpnUENLQ1l4WStFZnFPYTdn\nWmVKQVY3MTYzaVMrRjBJazFsZTdcXG5Gd2tlZ1NoVnpidkFoTEhLRG5jR1RKVjhiNXpUTitkZkR5\nYU4yMitwZkdDbUlLUG50dDJuay9QWDFCUkFCdWZ5XFxueGUvTU53T1ptcnVROUFLaDRYUlliVDFr\nNi9CT2NsM010OHJrc1lSRkVRS0JnUUQ3c2FPU1dYUE5HK09lb3dac1xcbjJObWVaMVBKRlpFMnpr\ndnFmenkwWXNTelFWYVFwSXJVSVBIaHQxL29TZmh6QmRaZXFjTjVSWTZ0bzRrZ29pK2hcXG5pTzJj\nckJrRkZjcUZWeGdrdm5RcjlaSXhaZ3RwWjFaZnI5dTc5QXQwdDh6M2NQNjJPcllneGZkZGhKL3Ba\nTjNLXFxuNUk1MGZQVC8xN054SjhJQTlMRUp3b2RXelFLQmdRRDU4ZnF4NGdDQVEwVWU0dyszcE5L\nMEpMOWhla2laK0k2VVxcbmVYZ0tUNGlrbmR2VkVwSVYwQlRFMFBXU3YyNnJ0bWowbDVwK1RFRERw\nd0hNRk9pTXlicm1uNXBuTU96UkF5a0VcXG5YZEhiRGFPZ0FtVGJPQ3ZuUTdqTmtyU0xYTXNaczls\nTzZ0c2dtRzBkZVZ3RGJqaEJsWWNUZUN5VFYxMVBURlFUXFxuMDJSeTZrOUVwd0tCZ0UzMlBFQjh6\nTTNmc2FYVndZdTlyOHJOSHJyT24yaG1oa1ZnbkNFVDV2SFBiMHptYXNwQ1xcbldpb2NidXk0M09o\na0NuN1AvdWgxanpoaHkzdjhRTnk4V1QyS1lVV2Q5bGxQSFA5a3J2OWUxYVhQZ2dGb0xPeitcXG53\nQVY4Rm9CVTNueFcxODhDYUovR01sVVpXNThqeXorcGhDYTQvZnF4aWJlbkpRUyt4b2ZMWmJpUkFv\nR0JBT0xlXFxuTHkwUzlGMTNzZCt0UGZEMDNJUEM3eXV0Y1FUQXhib2kzMENNbkh5L0JIRE1vR2pJ\nTEhIUk1YWDM2SjVYdmNCMVxcbmN6ZThRdlRVUEI5Znd3MDNkandyRmwwZjNYU0NKOUxjemNURkJv\nWVFaamROTHh5RkxkTDZuSUg2d2ljY1JkMEpcXG44OGdNVXM1ei8zN2xwbzV4Q09BMUxsQzI5ZTBx\nWFE1NXpDMDZYS1QxQW9HQU9mQ1NpeVVReDRFUmFtRzZCODcwXFxuaXFzbmcxUFhiemc2ZXYrNVo4\nZlY4MEFGRzFicmRJOWlSdGVMSk9jcGFBZ2RkSFpYcEtYMHZWS3JBMk9vYlI1YlxcblgyNEl6YTR4\nTlo4WVFBTGVWQ3hKNzFleUVyN0M2STRkQmFiZVcvT0tKTGFselFLUFJsYzRCT25YbVMvUHJ0dWNc\nXG5QUGVjcGhGeVdoak9TVU56RTZFa1MyRT1cXG4tLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tXFxu\nXCIsXG4gXCJjbGllbnRfZW1haWxcIjogXCJkZXYtNTk0QGFjaGlsaW8tZGV2LmlhbS5nc2Vydmlj\nZWFjY291bnQuY29tXCIsXG4gXCJjbGllbnRfaWRcIjogXCIxMDkyNjA0Mjk4NTg1NDE0NTM0NjNc\nIixcbiBcImF1dGhfdXJpXCI6IFwiaHR0cHM6Ly9hY2NvdW50cy5nb29nbGUuY29tL28vb2F1dGgy\nL2F1dGhcIixcbiBcInRva2VuX3VyaVwiOiBcImh0dHBzOi8vb2F1dGgyLmdvb2dsZWFwaXMuY29t\nL3Rva2VuXCIsXG4gXCJhdXRoX3Byb3ZpZGVyX3g1MDlfY2VydF91cmxcIjogXCJodHRwczovL3d3\ndy5nb29nbGVhcGlzLmNvbS9vYXV0aDIvdjEvY2VydHNcIixcbiBcImNsaWVudF94NTA5X2NlcnRf\ndXJsXCI6IFwiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vcm9ib3QvdjEvbWV0YWRhdGEveDUw\nOS9kZXYtNTk0JTQwYWNoaWxpby1kZXYuaWFtLmdzZXJ2aWNlYWNjb3VudC5jb21cIlxufVxuIiwi\ncXVlcmllcyI6W3siZGF0YXNldE5hbWUiOiJueWNfdHJpcHMiLCJzdGF0ZW1lbnQiOiJTRUxFQ1Qg\ncGF5bWVudF90eXBlIEFTIGFfNDk2OTMyMzk3LCBwYXNzZW5nZXJfY291bnQgQVMgYV8yMzcyMTkx\nMTQsIHJhdGVfY29kZSBBUyBhXzIxNjA4MTkwMCwgQ09VTlQoKikgQVMgYV80MTA1NzY5MjAgRlJP\nTSBgYWNoaWxpby1kZXZgLmBueWNfdHJpcHNgLmBueWNfdHJpcHNfc21hbGxfMjAxNWAgR1JPVVAg\nQlkgYV80OTY5MzIzOTcsIGFfMjM3MjE5MTE0LCBhXzIxNjA4MTkwMCIsIm1tdk5hbWUiOiJueWNf\ndHJpcHNfc21hbGxfMjAxNV9hY2hpbGlvXzE2NjMwMjI3MjkifSx7ImRhdGFzZXROYW1lIjoibnlj\nX3RyaXBzIiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSBBUyBhXzQ5NjkzMjM5Nywg\ncGFzc2VuZ2VyX2NvdW50IEFTIGFfMjM3MjE5MTE0LCBUSU1FU1RBTVBfVFJVTkMocGlja3VwX2Rh\ndGV0aW1lLCBXRUVLKSBBUyBhXzY1MjAzNzE0NCwgU1VNKHRpcF9hbW91bnQpIEFTIGFfMjM4OTA5\nNTIwIEZST00gYGFjaGlsaW8tZGV2YC5gbnljX3RyaXBzYC5gbnljX3RyaXBzX3NtYWxsXzIwMTVg\nIEdST1VQIEJZIGFfNDk2OTMyMzk3LCBhXzIzNzIxOTExNCwgYV82NTIwMzcxNDQiLCJtbXZOYW1l\nIjoibnljX3RyaXBzX3NtYWxsXzIwMTVfYWNoaWxpb18xNTQ0MjI0MTMifSx7ImRhdGFzZXROYW1l\nIjoibnljX3RyaXBzIiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSBBUyBhXzQ5Njkz\nMjM5NywgcGFzc2VuZ2VyX2NvdW50IEFTIGFfMjM3MjE5MTE0LCByYXRlX2NvZGUgQVMgYV8yMTYw\nODE5MDAsIEFWRyhmYXJlX2Ftb3VudCkgQVMgYV8yMDQ2Njk4MTAyIEZST00gYGFjaGlsaW8tZGV2\nYC5gbnljX3RyaXBzYC5gbnljX3RyaXBzX3NtYWxsXzIwMTVgIEdST1VQIEJZIGFfNDk2OTMyMzk3\nLCBhXzIzNzIxOTExNCwgYV8yMTYwODE5MDAiLCJtbXZOYW1lIjoibnljX3RyaXBzX3NtYWxsXzIw\nMTVfYWNoaWxpb18xMzIzNzg4Nzg1In0seyJkYXRhc2V0TmFtZSI6Im55Y190cmlwcyIsInN0YXRl\nbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUgQVMgYV80OTY5MzIzOTcsIHBhc3Nlbmdlcl9jb3Vu\ndCBBUyBhXzIzNzIxOTExNCwgQVZHKGZhcmVfYW1vdW50KSBBUyBhXzIwNDY2OTgxMDIgRlJPTSBg\nYWNoaWxpby1kZXZgLmBueWNfdHJpcHNgLmBueWNfdHJpcHNfc21hbGxfMjAxNWAgR1JPVVAgQlkg\nYV80OTY5MzIzOTcsIGFfMjM3MjE5MTE0IiwibW12TmFtZSI6Im55Y190cmlwc19zbWFsbF8yMDE1\nX2FjaGlsaW9fMTc0MDYxMjE1In0seyJkYXRhc2V0TmFtZSI6Im55Y190cmlwcyIsInN0YXRlbWVu\ndCI6IlNFTEVDVCBwYXltZW50X3R5cGUgQVMgYV80OTY5MzIzOTcsIHBhc3Nlbmdlcl9jb3VudCBB\nUyBhXzIzNzIxOTExNCwgcmF0ZV9jb2RlIEFTIGFfMjE2MDgxOTAwLCB2ZW5kb3JfaWQgQVMgYV8x\nMDg1OTcwNTc0LCBTVU0odG90YWxfYW1vdW50KSBBUyBhXzEyNzUwMTM5OTEgRlJPTSBgYWNoaWxp\nby1kZXZgLmBueWNfdHJpcHNgLmBueWNfdHJpcHNfc21hbGxfMjAxNWAgR1JPVVAgQlkgYV80OTY5\nMzIzOTcsIGFfMjM3MjE5MTE0LCBhXzIxNjA4MTkwMCwgYV8xMDg1OTcwNTc0IiwibW12TmFtZSI6\nIm55Y190cmlwc19zbWFsbF8yMDE1X2FjaGlsaW9fNTYyMTA4MjE0In0seyJkYXRhc2V0TmFtZSI6\nIm55Y190cmlwcyIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUgQVMgYV80OTY5MzIz\nOTcsIHBhc3Nlbmdlcl9jb3VudCBBUyBhXzIzNzIxOTExNCwgcmF0ZV9jb2RlIEFTIGFfMjE2MDgx\nOTAwLCBUSU1FU1RBTVBfVFJVTkMocGlja3VwX2RhdGV0aW1lLCBEQVkpIEFTIGFfMTU0NTYyMzM5\nOCwgQ09VTlQoKikgQVMgYV80MTA1NzY5MjAgRlJPTSBgYWNoaWxpby1kZXZgLmBueWNfdHJpcHNg\nLmBueWNfdHJpcHNfc21hbGxfMjAxNWAgR1JPVVAgQlkgYV80OTY5MzIzOTcsIGFfMjM3MjE5MTE0\nLCBhXzIxNjA4MTkwMCwgYV8xNTQ1NjIzMzk4IiwibW12TmFtZSI6Im55Y190cmlwc19zbWFsbF8y\nMDE1X2FjaGlsaW9fMjA1ODQyNTE3MyJ9LHsiZGF0YXNldE5hbWUiOiJueWNfdHJpcHMiLCJzdGF0\nZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlIEFTIGFfNDk2OTMyMzk3LCBwYXNzZW5nZXJfY291\nbnQgQVMgYV8yMzcyMTkxMTQsIHJhdGVfY29kZSBBUyBhXzIxNjA4MTkwMCwgVElNRVNUQU1QX1RS\nVU5DKHBpY2t1cF9kYXRldGltZSwgV0VFSykgQVMgYV82NTIwMzcxNDQsIFNVTSh0aXBfYW1vdW50\nKSBBUyBhXzIzODkwOTUyMCBGUk9NIGBhY2hpbGlvLWRldmAuYG55Y190cmlwc2AuYG55Y190cmlw\nc19zbWFsbF8yMDE1YCBHUk9VUCBCWSBhXzQ5NjkzMjM5NywgYV8yMzcyMTkxMTQsIGFfMjE2MDgx\nOTAwLCBhXzY1MjAzNzE0NCIsIm1tdk5hbWUiOiJueWNfdHJpcHNfc21hbGxfMjAxNV9hY2hpbGlv\nXzE4NDMxMzUxMjEifSx7ImRhdGFzZXROYW1lIjoibnljX3RyaXBzIiwic3RhdGVtZW50IjoiU0VM\nRUNUIHBheW1lbnRfdHlwZSBBUyBhXzQ5NjkzMjM5NywgcGFzc2VuZ2VyX2NvdW50IEFTIGFfMjM3\nMjE5MTE0LCBDT1VOVCgqKSBBUyBhXzQxMDU3NjkyMCBGUk9NIGBhY2hpbGlvLWRldmAuYG55Y190\ncmlwc2AuYG55Y190cmlwc19zbWFsbF8yMDE1YCBHUk9VUCBCWSBhXzQ5NjkzMjM5NywgYV8yMzcy\nMTkxMTQiLCJtbXZOYW1lIjoibnljX3RyaXBzX3NtYWxsXzIwMTVfYWNoaWxpb184MTQzOTkifSx7\nImRhdGFzZXROYW1lIjoibnljX3RyaXBzIiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlw\nZSBBUyBhXzQ5NjkzMjM5NywgcGFzc2VuZ2VyX2NvdW50IEFTIGFfMjM3MjE5MTE0LCB2ZW5kb3Jf\naWQgQVMgYV8xMDg1OTcwNTc0LCBTVU0odG90YWxfYW1vdW50KSBBUyBhXzEyNzUwMTM5OTEgRlJP\nTSBgYWNoaWxpby1kZXZgLmBueWNfdHJpcHNgLmBueWNfdHJpcHNfc21hbGxfMjAxNWAgR1JPVVAg\nQlkgYV80OTY5MzIzOTcsIGFfMjM3MjE5MTE0LCBhXzEwODU5NzA1NzQiLCJtbXZOYW1lIjoibnlj\nX3RyaXBzX3NtYWxsXzIwMTVfYWNoaWxpb18xNTA1MDE1ODEyIn0seyJkYXRhc2V0TmFtZSI6Im55\nY190cmlwcyIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUgQVMgYV80OTY5MzIzOTcs\nIHBhc3Nlbmdlcl9jb3VudCBBUyBhXzIzNzIxOTExNCwgVElNRVNUQU1QX1RSVU5DKHBpY2t1cF9k\nYXRldGltZSwgREFZKSBBUyBhXzE1NDU2MjMzOTgsIENPVU5UKCopIEFTIGFfNDEwNTc2OTIwIEZS\nT00gYGFjaGlsaW8tZGV2YC5gbnljX3RyaXBzYC5gbnljX3RyaXBzX3NtYWxsXzIwMTVgIEdST1VQ\nIEJZIGFfNDk2OTMyMzk3LCBhXzIzNzIxOTExNCwgYV8xNTQ1NjIzMzk4IiwibW12TmFtZSI6Im55\nY190cmlwc19zbWFsbF8yMDE1X2FjaGlsaW9fMTU2NzIzNjA0MyJ9XX0K",
			"messageId": "2070443601311540",
			"message_id": "2070443601311540",
			"publishTime": "2021-02-26T19:13:55.749Z",
			"publish_time": "2021-02-26T19:13:55.749Z"
		},
		"subscription": "projects/myproject/subscriptions/mysubscription"
	}`
	i2 := []byte(i1)
	var r1 PubSubMessage
	if err := json.Unmarshal(i2, &r1); err != nil {
		log.Fatalf("json.Unmarshal: should not be in error: %v\n", err)
		return
	}
}
