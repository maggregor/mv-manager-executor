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
			"data": "ewogICAgInNlcnZpY2VBY2NvdW50IjogIntcblwidHlwZVwiOiBcInNlcnZpY2VfYWNjb3VudFwiLFxuXCJwcm9qZWN0X2lkXCI6IFwiYWNoaWxpby1kZXZcIixcInByaXZhdGVfa2V5X2lkXCI6IFwiM2U5M2Y4M2M4MDcxNWVlNjc3ZDMyODI1NThmZjVhNTQwMGE4ZDg3MlwiLFwicHJpdmF0ZV9rZXlcIjogXCItLS0tLUJFR0lOIFBSSVZBVEUgS0VZLS0tLS1cbk1JSUV2Z0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktnd2dnU2tBZ0VBQW9JQkFRQ1VKSkdEQ0JxU1hINGdcblJTYWx2czZvZlIzbE54c3ExRVhGM2xPN0ZVU0taTHp6WmZqU1UzUllhamZocUVFTXhZNk9sZDdoYVVFS2hhQUtcbk0vVlpvWW5zZEhrdWVaeDdxNkNCdlpNdU03K0U4Y1UzVUJrSnIrK2dNQmJXdEozYzh6NEJ1UGRNazhnVWpNSExcbkhTSDRxQXNNRGxBM2RSMDhzZ3RCc0U0ZEx2TEkwS0E2dHBEYUc4dmxSK3pGcDBXcDdhT1BwZjZ2bGZTa3ZSemJcbmF1VWJhditFcWUvVVNGTGxkQWQ5dVMwUnlLQko1dVRDQ1hEcDdwNXBDRE1aUWZRN1pMM01ENzd1cVZ2Ui9qeVRcbmw3TGdJWms0MXE2MjdqTHpqeVdVV3ZQeUZZUXF6c2ExenhXdGlmQ29ERWZZelFBVWpFZ0IyZG9mcElXbHAxTEJcbjBPWWtFQzBqQWdNQkFBRUNnZ0VBUjV1SGhMVmhscmJEdHBpR2JsNlZhU2NWVU1zNnhCQ1lhMlVreHcwb05OY3RcbjZGODFqNEVNZjRKVVlhemhTS21qMTNORktXTWxyODdZaGFZSTVocnAzdFNNRVBWeXVPckVhb0kyeHdIM2hPMFNcbnZjSm5YK0hkcStvaTJBUzFPK3lOMytwSEM0cHhqOUZjQ1hDYkQ2VTJ0Y3lqWkxNay8wWTJuN2wzZ0s3RloyQ0lcbmlQbGswLzM5bThiN2dOZUhQNk8vSENzOUZ6d2x5amMyd0t2QU5hQnZnVmpoY2pSYUIrZHFKd0d4dzVkMUNURjhcbjRreEtJN2NwVE5mQmorenZIR2dQY3RxT0VOb1ZtTXRoQUVvNmFwY1RvOG9SRm9WSk9EUTVzaHFoMEdEK2pDbWxcbm0wVzgxR2xTK2Y1NmtQK0h2WnJiSUZuYXAzZUNiNFhoY1cwdFQ1YVdBUUtCZ1FESFo5ZmhXdGRzZlJkd0tZM21cbmJ0eGIvMzFVaVYvakl0RzhxeGd5aUYydFU0SDdhRXFWcTJRZ2RzMU1nN3NYRi9MNHU2VGNBaHRXR05EUFB0VGZcbllCM1gwRFBHVGxFQWkvQlhEWDZJNCtndUVrVjNLTGhOTE50MFl4VUowNnRuajRDNXNvYzRMYnJPWXZwbXJkZWVcbmpHWTdEYk9HZGlyUmZkd2dOVnlnQW1Lakl3S0JnUUMrTUNPMmJvVXc2L3djUENKOG9mNWVGeEpPVWFFbjlVbFhcbkRNUUpoMjRid1R5T0xwVjJMTWJ1MFdTUVF3eXZ0WGhQdTRETEthTFgvQmxLM0ZWU2piNmZWS2E4ckRrbzFmd0pcbjdxaXlSN0RIcVVJMlZhU3FIT0pjb2R5VXRuQjlCN2gzUDdGamV1c3paQ3FuNzk5Q3hoTHk1djVpcnE2eWFlZEZcbms4NmVwNkh1QVFLQmdRQ3NsQ015MXZ3NjJNakFPamFsaUswT3NrQTJPZTdURmYrUDJrUkc5OTZiYy9xN0s3ZmhcbmNVUm5GR3I2SGNoK1pDdlZPaEdrbCs0d3hhbHl0RjVvTlRlZTRJTHV6amtzei9CaVp6dHNIbE9FREN3eXhQczlcblQrVmZCYUhmcjJKWmJzbHA3aWs1WmRxWFpQSnlpMkpoeGdGMGVwYzFlRy8yZjRkVUdOcmhZUzhkSVFLQmdEdGlcbkxwUjAxOU9acTROcFoxOUErWHFKZTZiR3FDb2cxWUIvdE9wQXhpbVdZWUVIbkpWa2o4emJ4cW1ndGVKYktEREZcbkhBYlRma3d3SS9tUzZIVTVXdTdHOExUeWxYcGhyaXV2d0M3Q1libEY2UzJyK2hrL29aSkkyK1gzUFJKZENjVDZcbk56b1hsRDdjcE5FQS9kWU1vTGdGVmlLYjVtSkEzUHlGbFR0R25qSUJBb0dCQUxyVFA0STZHWEh5VTZWekxTN1pcbjRCdkdWQm9DajVtTk5aUmR0MGlTWVo4NkNBVFpLb25mV0MyY0dYQmdySytRekp2SndEdkFTd0M3MDRJcjZOczhcbnp0Y2J0Y1FlblNwN0dIalNPKzIxNGxMeDVDbzVzTkhWZHdQMjMxSVZMOFFYbG5Xbk44M3YwR1FaZkxrME9hb2dcbjZmSFJQMkZPcTMybExiWHRXeEw2NG5ST1xuLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLVwiLFwiY2xpZW50X2VtYWlsXCI6IFwiYWNoaWxpby1zYUBhY2hpbGlvLWRldi5pYW0uZ3NlcnZpY2VhY2NvdW50LmNvbVwiLFwiY2xpZW50X2lkXCI6IFwiMTEzNTI4NTk1NTcwMjI2MzM0NTUyXCIsXCJhdXRoX3VyaVwiOiBcImh0dHBzOi8vYWNjb3VudHMuZ29vZ2xlLmNvbS9vL29hdXRoMi9hdXRoXCIsXCJ0b2tlbl91cmlcIjogXCJodHRwczovL29hdXRoMi5nb29nbGVhcGlzLmNvbS90b2tlblwiLFwiYXV0aF9wcm92aWRlcl94NTA5X2NlcnRfdXJsXCI6IFwiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vb2F1dGgyL3YxL2NlcnRzXCIsXCJjbGllbnRfeDUwOV9jZXJ0X3VybFwiOiBcImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL3JvYm90L3YxL21ldGFkYXRhL3g1MDkvYWNoaWxpby1zYSU0MGFjaGlsaW8tZGV2LmlhbS5nc2VydmljZWFjY291bnQuY29tXCJ9IiwKICAgICJxdWVyaWVzIjogWwogICAgICAgIHsKICAgICAgICAgICAgImRhdGFzZXROYW1lIjogImV4ZWN1dG9yX2ludGVncmF0aW9uX3Rlc3QiLAogICAgICAgICAgICAibW12TmFtZSI6ICJteXRlc3RfbXZtIiwKICAgICAgICAgICAgInN0YXRlbWVudCI6ICJTRUxFQ1QgQ09VTlQoKikgQVMgYV80MTA1NzY5MjAgRlJPTSBgYWNoaWxpby10ZXN0YC5gZXhlY3V0b3JfaW50ZWdyYXRpb25fdGVzdGAuYGFfdGFibGVgIgogICAgICAgIH0KICAgIF0KfQ==",
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
	if r1.Message.Attributes.ServiceAccount.toString() != "{type:achilio-dev@achilio.com}" {
		log.Fatalf("Expected %s, got %s", "{type:achilio-dev@achilio.com}", r1.Message.Attributes.ServiceAccount)
	}
}
