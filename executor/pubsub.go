package executor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Attributes is the payload of the attributes field in the message of a Pub/Sub event.
// For the Terraform executor to work for Achilio's Terraview, it needs to follow this structure
type Attributes struct {
	AccessToken string   `json:"accessToken,omitempty"`
	ProjectID   string   `json:"projectId"`
	RegionID    string   `json:"regionId"`
	DatasetID   string   `json:"datasetId"`
	CmdType     string   `json:"cmdType"`
	Queries     []string `json:"queries"`
}

// Message is the payload of the message field of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type Message struct {
	Data       []byte     `json:"data,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
	ID         string     `json:"id"`
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Message      Message `json:"message"`
	Subscription string  `json:"subscription"`
}

// PubSub receives and processes a Pub/Sub push message.
func PubSub(w http.ResponseWriter, r *http.Request) {
	var m PubSubMessage
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	t := Terraform{m.Message, m.Message.Attributes, nil}
	err = t.Execute()
}
