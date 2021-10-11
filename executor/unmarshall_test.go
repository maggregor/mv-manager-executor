package executor

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAttributeUnmarshall1(t *testing.T) {
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "workspace"
			},
			"data": "SGVsbG8gQ2xvdWQgUHViL1N1YiEgSGVyZSBpcyBteSBtZXNzYWdlIQ==",
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
		log.Fatalln("json.Unmarshal: should be in error")
		return
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshall2(t *testing.T) {
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "workspace",
				"projectId": "myproject"
			},
			"data": "SGVsbG8gQ2xvdWQgUHViL1N1YiEgSGVyZSBpcyBteSBtZXNzYWdlIQ==",
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

func TestAttributeUnmarshall3(t *testing.T) {
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject"
			},
			"data": "SGVsbG8gQ2xvdWQgUHViL1N1YiEgSGVyZSBpcyBteSBtZXNzYWdlIQ==",
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
		log.Fatalln("json.Unmarshal: should be in error")
		return
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshall4(t *testing.T) {
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject",
				"regionId": "my-region"
			},
			"data": "SGVsbG8gQ2xvdWQgUHViL1N1YiEgSGVyZSBpcyBteSBtZXNzYWdlIQ==",
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
		log.Fatalln("json.Unmarshal: should be in error")
		return
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshall5(t *testing.T) {
	// data is: SELECT 1;SELECT 2
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject",
				"regionId": "my-region",
				"datasetName": "mydataset"
			},
			"data": "U0VMRUNUIDE7U0VMRUNUIDI=",
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
		log.Fatalf("json.Unmarshal: should not be in error %v\n", r1)
		return
	} else if len(r1.Message.Attributes.Queries) != 2 {
		log.Fatalf("Query len is %v, expected 2\n", len(r1.Message.Attributes.Queries))
	} else if r1.Message.Attributes.Queries[0] != "SELECT 1" {
		log.Fatalf("Query 1 is %v expected 'SELECT 1'\n", r1.Message.Attributes.Queries[0])
	} else if r1.Message.Attributes.Queries[1] != "SELECT 2" {
		log.Fatalf("Query 2 is %v, expected 'SELECT 2'\n", r1.Message.Attributes.Queries[1])
	} else if r1.Message.Attributes.CmdType != "apply" {
		log.Fatalf("cmdType is %v, expected 'apply'\n", r1.Message.Attributes.CmdType)
	} else if r1.Message.Attributes.ProjectID != "myproject" {
		log.Fatalf("Attribute projectId is %v, expected 'myproject'", r1.Message.Attributes.ProjectID)
	} else if r1.Message.Attributes.RegionID != "my-region" {
		log.Fatalf("Attribute regionId is %v, expected 'my-region'", r1.Message.Attributes.RegionID)
	} else if r1.Message.Attributes.DatasetName != "mydataset" {
		log.Fatalf("Attribute datasetName is %v, expected 'mydataset'", r1.Message.Attributes.DatasetName)
	} else if r1.Message.Attributes.AccessToken != "value" {
		log.Fatalf("Attribute accessToken is %v, expected 'value'", r1.Message.Attributes.AccessToken)
	} else {
		log.Printf("%v\n", r1.Message.Attributes)
	}
}

func TestAttributeUnmarshall6(t *testing.T) {
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "unvalidcmd",
				"projectId": "myproject"
			},
			"data": "SGVsbG8gQ2xvdWQgUHViL1N1YiEgSGVyZSBpcyBteSBtZXNzYWdlIQ==",
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
		log.Fatalf("json.Unmarshal: should be in error %v\n", r1)
		return
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshallMessageData(t *testing.T) {
	// data is: SELECT 1;
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject",
				"regionId": "myregion",
				"datasetName": "mydataset"
			},
			"data": "U0VMRUNUIDE7",
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
		log.Fatalf("json.Unmarshal: should not be in error %v\n", r1)
		return
	} else if len(r1.Message.Attributes.Queries) != 1 {
		log.Fatalf("Queries length is %v, expected 1.\n", len(r1.Message.Attributes.Queries))
	} else if r1.Message.Attributes.Queries[0] != "SELECT 1" {
		log.Fatalf("Query is %v, expected 'SELECT 1'", r1.Message.Attributes.Queries[0])
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshallMessageDataEmpty(t *testing.T) {
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject",
				"regionId": "myregion",
				"datasetName": "mydataset"
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
	if err := json.Unmarshal(i2, &r1); err != nil {
		log.Fatalf("json.Unmarshal: should not be in error %v\n", r1)
		return
	} else if len(r1.Message.Attributes.Queries) != 0 {
		log.Fatalf("Queries length is %v, expected 0. Empty query should be removed\n", len(r1.Message.Attributes.Queries))
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshallMessageDataVeryLong(t *testing.T) {
	// data is: SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject",
				"regionId": "myregion",
				"datasetName": "mydataset"
			},
			"data": "U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZTtTRUxFQ1QgcGF5bWVudF90eXBlLCBTVU0odG90YWxfYW1vdW50KSBhcyBjb2wxIEZST00gYWNoaWxpby1kZXYubnljX3RyaXBzLnRsY195ZWxsb3dfdHJpcHNfMjAxNV9zbWFsbCBHUk9VUCBCWSBwYXltZW50X3R5cGU7U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBlOw==",
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
		log.Fatalf("json.Unmarshal: should not be in error %v\n", r1)
		return
	} else if len(r1.Message.Attributes.Queries) != 40 {
		log.Fatalf("Queries length is %v, expected 40.\n", len(r1.Message.Attributes.Queries))
	} else {
		log.Printf("%v\n", err)
	}
}

func TestAttributeUnmarshallEscapedQuotesInData(t *testing.T) {
	// data is: SELECT (FORMAT_DATETIME(\"%B\", DATETIME(pickup_datetime))) IN (\"December\", \"January\", \"June\", \"March\") AS a_386307744
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject",
				"regionId": "myregion",
				"datasetName": "mydataset"
			},
			"data": "U0VMRUNUIChGT1JNQVRfREFURVRJTUUoXCIlQlwiLCBEQVRFVElNRShwaWNrdXBfZGF0ZXRpbWUpKSkgSU4gKFwiRGVjZW1iZXJcIiwgXCJKYW51YXJ5XCIsIFwiSnVuZVwiLCBcIk1hcmNoXCIpIEFTIGFfMzg2MzA3NzQ0",
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
		log.Fatalf("json.Unmarshal: should not be in error %v\n", r1)
	}
	log.Printf("Query is %v", r1.Message.Attributes.Queries[0])
	if r1.Message.Attributes.Queries[0] != `SELECT (FORMAT_DATETIME("%B", DATETIME(pickup_datetime))) IN ("December", "January", "June", "March") AS a_386307744` {
		log.Fatalf("Query is %v, expected %v", r1.Message.Attributes.Queries[0], `SELECT (FORMAT_DATETIME("%B", DATETIME(pickup_datetime))) IN ("December", "January", "June", "March") AS a_386307744`)
	}
}
