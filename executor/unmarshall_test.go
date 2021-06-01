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
			"data": "SELECT 1;SELECT 2",
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
	} else if r1.Message.Attributes.Queries[0] != "SELECT 1" {
		log.Fatalf("Query 1 is %v expected 'SELECT 1'\n", r1.Message.Attributes.Queries[0])
	} else if r1.Message.Attributes.Queries[1] != "SELECT 2" {
		log.Fatalf("Query 2 is %v, expected 'SELECT 2'\n", r1.Message.Attributes.Queries[0])
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
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject",
				"regionId": "myregion",
				"datasetId": "mydataset"
			},
			"data": "SELECT 1;",
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
	} else if len(r1.Message.Attributes.Queries) > 1 {
		log.Fatalf("Queries length is %v, expected 1. Empty query should be removed\n", len(r1.Message.Attributes.Queries))
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
				"datasetId": "mydataset"
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
	i1 := `{
		"message": {
			"attributes": {
				"accessToken": "value",
				"cmdType": "apply",
				"projectId": "myproject",
				"regionId": "myregion",
				"datasetId": "mydataset"
			},
			"data": "SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type;",
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
		log.Fatalf("Queries length is %v, expected 40. Empty query should be removed\n", len(r1.Message.Attributes.Queries))
	} else {
		log.Printf("%v\n", err)
	}
}
