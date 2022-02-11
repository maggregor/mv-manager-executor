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
	// data is: [{"mydataset1":"SELECT 1"}]
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W3sibXlkYXRhc2V0MSI6IlNFTEVDVCAxIn1d",
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
	// data is: [{"mydataset1":"SELECT 1","mydataset2":"SELECT 2"]
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W3sibXlkYXRhc2V0MSI6IlNFTEVDVCAxIn0seyJteWRhdGFzZXQyIjoiU0VMRUNUIDIifV0=",
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
	// data is: [{"mydataset1":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset2":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset3":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset4":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset5":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset6":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset7":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset8":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset9":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset10":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset11":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset12":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset13":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset14":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset15":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset16":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset17":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset18":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset19":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset20":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset21":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset22":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset23":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset24":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset25":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset26":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset27":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset28":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset29":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset30":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset31":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset32":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset33":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset34":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset35":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset36":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset37":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset38":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset39":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"},{"mydataset40":"SELECTpayment_type,SUM(total_amount)ascol1FROMachilio-dev.nyc_trips.tlc_yellow_trips_2015_smallGROUPBYpayment_type"}]
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W3sibXlkYXRhc2V0MSI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MiI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MyI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0NCI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0NSI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0NiI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0NyI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0OCI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0OSI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MTAiOiJTRUxFQ1RwYXltZW50X3R5cGUsU1VNKHRvdGFsX2Ftb3VudClhc2NvbDFGUk9NYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbEdST1VQQllwYXltZW50X3R5cGUifSx7Im15ZGF0YXNldDExIjoiU0VMRUNUcGF5bWVudF90eXBlLFNVTSh0b3RhbF9hbW91bnQpYXNjb2wxRlJPTWFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGxHUk9VUEJZcGF5bWVudF90eXBlIn0seyJteWRhdGFzZXQxMiI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MTMiOiJTRUxFQ1RwYXltZW50X3R5cGUsU1VNKHRvdGFsX2Ftb3VudClhc2NvbDFGUk9NYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbEdST1VQQllwYXltZW50X3R5cGUifSx7Im15ZGF0YXNldDE0IjoiU0VMRUNUcGF5bWVudF90eXBlLFNVTSh0b3RhbF9hbW91bnQpYXNjb2wxRlJPTWFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGxHUk9VUEJZcGF5bWVudF90eXBlIn0seyJteWRhdGFzZXQxNSI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MTYiOiJTRUxFQ1RwYXltZW50X3R5cGUsU1VNKHRvdGFsX2Ftb3VudClhc2NvbDFGUk9NYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbEdST1VQQllwYXltZW50X3R5cGUifSx7Im15ZGF0YXNldDE3IjoiU0VMRUNUcGF5bWVudF90eXBlLFNVTSh0b3RhbF9hbW91bnQpYXNjb2wxRlJPTWFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGxHUk9VUEJZcGF5bWVudF90eXBlIn0seyJteWRhdGFzZXQxOCI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MTkiOiJTRUxFQ1RwYXltZW50X3R5cGUsU1VNKHRvdGFsX2Ftb3VudClhc2NvbDFGUk9NYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbEdST1VQQllwYXltZW50X3R5cGUifSx7Im15ZGF0YXNldDIwIjoiU0VMRUNUcGF5bWVudF90eXBlLFNVTSh0b3RhbF9hbW91bnQpYXNjb2wxRlJPTWFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGxHUk9VUEJZcGF5bWVudF90eXBlIn0seyJteWRhdGFzZXQyMSI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MjIiOiJTRUxFQ1RwYXltZW50X3R5cGUsU1VNKHRvdGFsX2Ftb3VudClhc2NvbDFGUk9NYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbEdST1VQQllwYXltZW50X3R5cGUifSx7Im15ZGF0YXNldDIzIjoiU0VMRUNUcGF5bWVudF90eXBlLFNVTSh0b3RhbF9hbW91bnQpYXNjb2wxRlJPTWFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGxHUk9VUEJZcGF5bWVudF90eXBlIn0seyJteWRhdGFzZXQyNCI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MjUiOiJTRUxFQ1RwYXltZW50X3R5cGUsU1VNKHRvdGFsX2Ftb3VudClhc2NvbDFGUk9NYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbEdST1VQQllwYXltZW50X3R5cGUifSx7Im15ZGF0YXNldDI2IjoiU0VMRUNUcGF5bWVudF90eXBlLFNVTSh0b3RhbF9hbW91bnQpYXNjb2wxRlJPTWFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGxHUk9VUEJZcGF5bWVudF90eXBlIn0seyJteWRhdGFzZXQyNyI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MjgiOiJTRUxFQ1RwYXltZW50X3R5cGUsU1VNKHRvdGFsX2Ftb3VudClhc2NvbDFGUk9NYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbEdST1VQQllwYXltZW50X3R5cGUifSx7Im15ZGF0YXNldDI5IjoiU0VMRUNUcGF5bWVudF90eXBlLFNVTSh0b3RhbF9hbW91bnQpYXNjb2wxRlJPTWFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGxHUk9VUEJZcGF5bWVudF90eXBlIn0seyJteWRhdGFzZXQzMCI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MzEiOiJTRUxFQ1RwYXltZW50X3R5cGUsU1VNKHRvdGFsX2Ftb3VudClhc2NvbDFGUk9NYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbEdST1VQQllwYXltZW50X3R5cGUifSx7Im15ZGF0YXNldDMyIjoiU0VMRUNUcGF5bWVudF90eXBlLFNVTSh0b3RhbF9hbW91bnQpYXNjb2wxRlJPTWFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGxHUk9VUEJZcGF5bWVudF90eXBlIn0seyJteWRhdGFzZXQzMyI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MzQiOiJTRUxFQ1RwYXltZW50X3R5cGUsU1VNKHRvdGFsX2Ftb3VudClhc2NvbDFGUk9NYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbEdST1VQQllwYXltZW50X3R5cGUifSx7Im15ZGF0YXNldDM1IjoiU0VMRUNUcGF5bWVudF90eXBlLFNVTSh0b3RhbF9hbW91bnQpYXNjb2wxRlJPTWFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGxHUk9VUEJZcGF5bWVudF90eXBlIn0seyJteWRhdGFzZXQzNiI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0MzciOiJTRUxFQ1RwYXltZW50X3R5cGUsU1VNKHRvdGFsX2Ftb3VudClhc2NvbDFGUk9NYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbEdST1VQQllwYXltZW50X3R5cGUifSx7Im15ZGF0YXNldDM4IjoiU0VMRUNUcGF5bWVudF90eXBlLFNVTSh0b3RhbF9hbW91bnQpYXNjb2wxRlJPTWFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGxHUk9VUEJZcGF5bWVudF90eXBlIn0seyJteWRhdGFzZXQzOSI6IlNFTEVDVHBheW1lbnRfdHlwZSxTVU0odG90YWxfYW1vdW50KWFzY29sMUZST01hY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsR1JPVVBCWXBheW1lbnRfdHlwZSJ9LHsibXlkYXRhc2V0NDAiOiJTRUxFQ1RwYXltZW50X3R5cGUsU1VNKHRvdGFsX2Ftb3VudClhc2NvbDFGUk9NYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbEdST1VQQllwYXltZW50X3R5cGUifV0=",
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
	// data is: [{"mydataset1":"SELECT (FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime))) IN (\"December\", \"January\", \"June\", \"March\") AS a_386307744"}]
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W3sibXlkYXRhc2V0MSI6IlNFTEVDVCAoRk9STUFUX0RBVEVUSU1FKFwiJUJcIiwgREFURVRJTUUocGlja3VwX2RhdGV0aW1lKSkpIElOIChcIkRlY2VtYmVyXCIsIFwiSmFudWFyeVwiLCBcIkp1bmVcIiwgXCJNYXJjaFwiKSBBUyBhXzM4NjMwNzc0NCJ9XQ==",
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
	// data is: [{"mydataset1":"SELECT 1"},{"mydataset2":"SELECT 2"},{"mydataset2":"SELECT 2"},{"mydataset2":"SELECT 1"}]
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "W3sibXlkYXRhc2V0MSI6IlNFTEVDVCAxIn0seyJteWRhdGFzZXQyIjoiU0VMRUNUIDIifSx7Im15ZGF0YXNldDIiOiJTRUxFQ1QgMiJ9LHsibXlkYXRhc2V0MiI6IlNFTEVDVCAxIn1d",
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
