package executor

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAttributeUnmarshallErrorNoProject(t *testing.T) {
	log.Println("TestAttributeUnmarshallErrorNoProject")
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "workspace"
			},
			"data": "W10=",
			"messageId": "2070443601311540",
			"message_id": "2070443601311540",
			"publishTime": "2021-02-26T19:13:55.749Z",
			"publish_time": "2021-02-26T19:13:55.749Z"
		},
	    "subscription": "projects/myproject/subscriptions/mysubscription"
	}`
	i2 := []byte(i1)
	var r1 PubSubMessage
	if err := json.Unmarshal(i2, &r1); err == nil {
		log.Fatalln("json.Unmarshal: should be in error: no project")
		return
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshallInvalidData(t *testing.T) {
	log.Println("TestAttributeUnmarshallInvalidData")
	// Data is: activating project: notre-vie
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "workspace",
				"projectId": "myproject"
			},
			"data": "YWN0aXZhdGluZyBwcm9qZWN0OiBub3RyZS12aWU=",
			"messageId": "2070443601311540",
			"message_id": "2070443601311540",
			"publishTime": "2021-02-26T19:13:55.749Z",
			"publish_time": "2021-02-26T19:13:55.749Z"
		},
		"subscription": "projects/myproject/subscriptions/mysubscription"
	}`
	i2 := []byte(i1)
	var r1 PubSubMessage
	if err := json.Unmarshal(i2, &r1); err == nil {
		log.Fatalln("json.Unmarshal: should be in error: data is not the correct format")
		return
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshallInvalidCmd1(t *testing.T) {
	log.Println("TestAttributeUnmarshallInvalidCmd1")
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "foo"
			},
			"data": "W10=",
			"messageId": "2070443601311540",
			"message_id": "2070443601311540",
			"publishTime": "2021-02-26T19:13:55.749Z",
			"publish_time": "2021-02-26T19:13:55.749Z"
		},
	    "subscription": "projects/myproject/subscriptions/mysubscription"
	}`
	i2 := []byte(i1)
	var r1 PubSubMessage
	if err := json.Unmarshal(i2, &r1); err == nil {
		log.Fatalf("json.Unmarshal: should be in error: invalid cmd, got %v", err)
		return
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshallInvalidCmd2(t *testing.T) {
	log.Println("TestAttributeUnmarshallInvalidCmd2")
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "unvalidcmd",
				"projectId": "myproject"
			},
			"data": "W10=",
			"messageId": "2070443601311540",
			"message_id": "2070443601311540",
			"publishTime": "2021-02-26T19:13:55.749Z",
			"publish_time": "2021-02-26T19:13:55.749Z"
		},
	    "subscription": "projects/myproject/subscriptions/mysubscription"
	}`
	i2 := []byte(i1)
	var r1 PubSubMessage
	if err := json.Unmarshal(i2, &r1); err == nil {
		log.Fatalf("json.Unmarshal: should be in error invalid cmd, got %v\n", err)
		return
	} else {
		log.Printf("%v\n", err)
	}
}

func TestUnmarshallError(t *testing.T) {
	log.Println("TestUnmarshallError")
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "workspace",
				"projectId": "myproject",
			},
			"data": "W10=",
			"messageId": "2070443601311540",
			"message_id": "2070443601311540",
			"publishTime": "2021-02-26T19:13:55.749Z",
			"publish_time": "2021-02-26T19:13:55.749Z"
		},
	    "subscription": "projects/myproject/subscriptions/mysubscription"
	}`
	i2 := []byte(i1)
	var r1 PubSubMessage
	if err := json.Unmarshal(i2, &r1); err == nil {
		log.Fatalf("json.Unmarshal: should in error: wrong json formatting: %v\n", err)
	}
}

func TestAttributeUnmarshallWorkspace(t *testing.T) {
	log.Println("TestAttributeUnmarshallWorkspace")
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "workspace",
				"projectId": "myproject"
			},
			"data": "W10=",
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
	} else if r1.Message.Attributes.CmdType != "workspace" {
		log.Fatalf("Attribute cmdType is %v, expected 'workspace'", r1.Message.Attributes.CmdType)
	} else if r1.Message.Attributes.ProjectID != "myproject" {
		log.Fatalf("Attribute projectId is %v, expected 'myproject'", r1.Message.Attributes.CmdType)
	} else if r1.Message.Attributes.AccessToken != "value" {
		log.Fatalf("Attribute accessToken is %v, expected 'value'", r1.Message.Attributes.CmdType)
	} else {
		log.Printf("%v\n", r1.Message.Attributes)
	}
}

func TestAttributeUnmarshallWorkspaceNoToken(t *testing.T) {
	log.Println("TestAttributeUnmarshallWorkspaceNoToken")
	i1 := `{
		"message": {
			"attributes": {
				"cmdType": "workspace",
				"projectId": "myproject"
			},
			"data": "W10=",
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
	} else if r1.Message.Attributes.CmdType != "workspace" {
		log.Fatalf("Attribute cmdType is %v, expected 'workspace'", r1.Message.Attributes.CmdType)
	} else if r1.Message.Attributes.ProjectID != "myproject" {
		log.Fatalf("Attribute projectId is %v, expected 'myproject'", r1.Message.Attributes.CmdType)
	} else {
		log.Printf("%v\n", r1.Message.Attributes)
	}
}

func TestAttributeUnmarshallEmptyQueries(t *testing.T) {
	log.Println("TestAttributeUnmarshallEmptyQueries")
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W10=",
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
		log.Fatalln("json.Unmarshal: should not be in error")
	}
	if len(r1.Message.Attributes.Queries) != 0 {
		log.Fatalf("There are %v queries, expected 0", len(r1.Message.Attributes.Queries))
	}
}

func TestAttributeUnmarshallValidData1(t *testing.T) {
	log.Println("TestAttributeUnmarshallValidData1")
	// data is: [{"datasetName":"mydataset","mmvName":"achilio_1234","statement":"SELECT 1"}]
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W3siZGF0YXNldE5hbWUiOiJteWRhdGFzZXQiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0Iiwic3RhdGVtZW50IjoiU0VMRUNUIDEifV0=",
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
		log.Fatalf("json.Unmarshal: should not be in error, got %v\n", err)
		return
	} else if len(r1.Message.Attributes.Queries) != 1 {
		log.Fatalf("Queries length is %v, expected 1.\n", len(r1.Message.Attributes.Queries))
	} else if r1.Message.Attributes.Queries[0].Statement != "SELECT 1" {
		log.Fatalf("Query is %v, expected 'SELECT 1'", r1.Message.Attributes.Queries[0])
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshallValidData2(t *testing.T) {
	log.Println("TestAttributeUnmarshallValidData2")
	// data is: [{"datasetName":"mydataset1","mmvName":"achilio_1234","statement":"SELECT 1"},{"datasetName":"mydataset2","mmvName":"achilio_12345","statement":"SELECT 2"}]
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W3siZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNCIsInN0YXRlbWVudCI6IlNFTEVDVCAxIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDIiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCAyIn1d",
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
		log.Fatalf("json.Unmarshal: should not be in error %v\n", err)
		return
	} else if len(r1.Message.Attributes.Queries) != 2 {
		log.Fatalf("Query len is %v, expected 2\n", len(r1.Message.Attributes.Queries))
	} else if r1.Message.Attributes.Queries[0].Statement != "SELECT 1" {
		log.Fatalf("Query 1 is %v expected 'SELECT 1'\n", r1.Message.Attributes.Queries[0].Statement)
	} else if r1.Message.Attributes.Queries[1].Statement != "SELECT 2" {
		log.Fatalf("Query 2 is %v, expected 'SELECT 2'\n", r1.Message.Attributes.Queries[1].Statement)
	} else if r1.Message.Attributes.Queries[0].DatasetName != "mydataset1" {
		log.Fatalf("Query 1 dataset is %v expected 'mydataset1'\n", r1.Message.Attributes.Queries[0].DatasetName)
	} else if r1.Message.Attributes.Queries[1].DatasetName != "mydataset2" {
		log.Fatalf("Query 2 dataset is %v, expected 'mydataset2'\n", r1.Message.Attributes.Queries[1].DatasetName)
	} else if r1.Message.Attributes.Queries[0].MmvName != "achilio_1234" {
		log.Fatalf("Query 1 name is %v expected 'achilio_1234'\n", r1.Message.Attributes.Queries[0].MmvName)
	} else if r1.Message.Attributes.Queries[1].MmvName != "achilio_12345" {
		log.Fatalf("Query 2 name is %v expected 'achilio_12345'\n", r1.Message.Attributes.Queries[1].MmvName)
	} else if r1.Message.Attributes.CmdType != "apply" {
		log.Fatalf("cmdType is %v, expected 'apply'\n", r1.Message.Attributes.CmdType)
	} else if r1.Message.Attributes.ProjectID != "myproject" {
		log.Fatalf("Attribute projectId is %v, expected 'myproject'", r1.Message.Attributes.ProjectID)
	} else if r1.Message.Attributes.AccessToken != "value" {
		log.Fatalf("Attribute accessToken is %v, expected 'value'", r1.Message.Attributes.AccessToken)
	} else {
		log.Printf("%v\n", r1.Message.Attributes)
	}
}

func TestAttributeUnmarshallMessageDataEmpty(t *testing.T) {
	log.Println("TestAttributeUnmarshallMessageDataEmpty")
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "",
			"messageId": "2070443601311540",
			"message_id": "2070443601311540",
			"publishTime": "2021-02-26T19:13:55.749Z",
			"publish_time": "2021-02-26T19:13:55.749Z"
		},
	    "subscription": "projects/myproject/subscriptions/mysubscription"
	}`
	i2 := []byte(i1)
	var r1 PubSubMessage
	if err := json.Unmarshal(i2, &r1); err == nil {
		log.Fatalf("json.Unmarshal: should be in error")
		return
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshallMessageDataEmptyArray(t *testing.T) {
	log.Println("TestAttributeUnmarshallMessageDataEmptyArray")
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W10=",
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
		log.Fatalf("json.Unmarshal: should not be in error %v\n", err)
		return
	} else if len(r1.Message.Attributes.Queries) != 0 {
		log.Fatalf("Queries length is %v, expected 0. Empty query should be removed\n", len(r1.Message.Attributes.Queries))
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshallMessageDataVeryLong(t *testing.T) {
	log.Println("TestAttributeUnmarshallMessageDataVeryLong")
	// data is: [{"datasetName":"mydataset1","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset2","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset3","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset4","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset5","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset6","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset7","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset8","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset9","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset10","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset11","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset12","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset13","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset14","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset15","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset16","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset17","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset18","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset19","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset20","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset21","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset22","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset23","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset24","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset25","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset26","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset27","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset28","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset29","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset30","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset31","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset32","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset33","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset34","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset35","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset36","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset37","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset38","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset39","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"datasetName":"mydataset40","mmvName":"achilio_12345","statement":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"}]
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W3siZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MiIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDMiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQ0IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0NSIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDYiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQ3IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0OCIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDkiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxMCIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDExIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MTIiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxMyIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDE0IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MTUiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxNiIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDE3IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MTgiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQxOSIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDIwIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MjEiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQyMiIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDIzIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MjQiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQyNSIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDI2IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MjciLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQyOCIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDI5IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MzAiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQzMSIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDMyIiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MzMiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQzNCIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDM1IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MzYiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQzNyIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn0seyJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDM4IiwibW12TmFtZSI6ImFjaGlsaW9fMTIzNDUiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGUifSx7ImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MzkiLCJtbXZOYW1lIjoiYWNoaWxpb18xMjM0NSIsInN0YXRlbWVudCI6IlNFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZSJ9LHsiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQ0MCIsIm1tdk5hbWUiOiJhY2hpbGlvXzEyMzQ1Iiwic3RhdGVtZW50IjoiU0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlIn1d",
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
		log.Fatalf("json.Unmarshal: should not be in error %v\n", err)
		return
	} else if len(r1.Message.Attributes.Queries) != 40 {
		log.Fatalf("Queries length is %v, expected 40.\n", len(r1.Message.Attributes.Queries))
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshallEscapedQuotesInData(t *testing.T) {
	log.Println("TestAttributeUnmarshallEscapedQuotesInData")
	// data is: [{"mmvName":"achilio_1234","datasetName":"mydataset","statement":"SELECT (FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime))) IN (\"December\", \"January\", \"June\", \"March\") AS a_386307744"}]
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W3sibW12TmFtZSI6ImFjaGlsaW9fMTIzNCIsImRhdGFzZXROYW1lIjoibXlkYXRhc2V0Iiwic3RhdGVtZW50IjoiU0VMRUNUIChGT1JNQVRfREFURVRJTUUoXCIlQlwiLCBEQVRFVElNRShwaWNrdXBfZGF0ZXRpbWUpKSkgSU4gKFwiRGVjZW1iZXJcIiwgXCJKYW51YXJ5XCIsIFwiSnVuZVwiLCBcIk1hcmNoXCIpIEFTIGFfMzg2MzA3NzQ0In1d",
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
		log.Fatalf("json.Unmarshal: should not be in error %v\n", err)
	}
	log.Printf("Query is %v", r1.Message.Attributes.Queries[0])
	if r1.Message.Attributes.Queries[0].Statement != `SELECT (FORMAT_DATETIME("%B", DATETIME(pickup_datetime))) IN ("December", "January", "June", "March") AS a_386307744` {
		log.Fatalf("Query is %v, expected %v", r1.Message.Attributes.Queries[0], `SELECT (FORMAT_DATETIME("%B", DATETIME(pickup_datetime))) IN ("December", "January", "June", "March") AS a_386307744`)
	}
}

func TestAttributeUnmarshallDuplicateQuery(t *testing.T) {
	log.Println("TestAttributeUnmarshallDuplicateQuery")
	// data is: [{"mmvName":"achilio_1234","datasetName":"mydataset1","statement":"SELECT1"},{"mmvName":"achilio_1234","datasetName":"mydataset1","statement":"SELECT2"},{"mmvName":"achilio_1234","datasetName":"mydataset2","statement":"SELECT1"},{"mmvName":"achilio_1234","datasetName":"mydataset1","statement":"SELECT1"}]
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W3sibW12TmFtZSI6ImFjaGlsaW9fMTIzNCIsImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MSIsInN0YXRlbWVudCI6IlNFTEVDVDEifSx7Im1tdk5hbWUiOiJhY2hpbGlvXzEyMzQiLCJkYXRhc2V0TmFtZSI6Im15ZGF0YXNldDEiLCJzdGF0ZW1lbnQiOiJTRUxFQ1QyIn0seyJtbXZOYW1lIjoiYWNoaWxpb18xMjM0IiwiZGF0YXNldE5hbWUiOiJteWRhdGFzZXQyIiwic3RhdGVtZW50IjoiU0VMRUNUMSJ9LHsibW12TmFtZSI6ImFjaGlsaW9fMTIzNCIsImRhdGFzZXROYW1lIjoibXlkYXRhc2V0MSIsInN0YXRlbWVudCI6IlNFTEVDVDEifV0=",
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
		log.Fatalf("json.Unmarshal: should not be in error %v\n", err)
	}
	if len(r1.Message.Attributes.Queries) != 3 {
		log.Fatalf("There are %v queries, expected 3", len(r1.Message.Attributes.Queries))
	}
}
