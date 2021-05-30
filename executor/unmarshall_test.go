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
				"regionId": "myregion"
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
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject",
				"regionId": "myregion",
				"datasetId": "mydataset",
				"queries": ["SELECT 1", "SELECT 2"]
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
		log.Fatalf("json.Unmarshal: should not be in error %v\n", r1)
		return
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

func TestAttributeUnmarshall7(t *testing.T) {
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject",
				"regionId": "myregion",
				"datasetId": "mydataset",
				"queries": "SELECT 1"
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
