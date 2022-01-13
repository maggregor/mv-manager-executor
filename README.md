# Achilio mv-manager-executor

[![codecov](https://codecov.io/gh/achilio/mv-manager-executor/branch/master/graph/badge.svg?token=OLM9U79QD4)](https://codecov.io/gh/achilio/mv-manager-executor)

## Roadmap

-   [x] Handling notification message to build the command to execute
-   [x] Add terraform module to actually execute

## Run locally

```
docker build -t local/mv-manager-executor .
docker run -it -p 8080:8080 -e AWS_ACCESS_KEY_ID=<AWS_ACCESS_KEY> -e AWS_SECRET_ACCESS_KEY=<AWS_SECRET_KEY> -p 8080 local/mv-manager-executor
```

## Send a notification

Locally

`./local_test.sh`

Remotely

`./remote_test.sh`

## Deploy on GCP

`gcloud builds submit --config cloudbuild.yaml`

Note: Some prequisites are necessary. Read the Cloud Run with Pub/Sub documentation for more info: https://cloud.google.com/run/docs/tutorials/pubsub#before-you-begin

## Send notifications with Pub/Sub

Documentation to send messages via Pub/Sub in Java (and many other languages): https://cloud.google.com/pubsub/docs/publisher

For the mv-manager-executor to work, the message must have the following attributes:

-   "cmdType": enum("apply", "workspace") // Apply is for creating a new set of views. Workspace is for creating a new workspace (new project activated)
-   "projectId": string
-   "accessToken: string
-   "datasetName": string // Not required if cmdType is workspace

The message data is a list of ; separated SELECT queries representing the views to create. The message is sent as a base64 encoded string when using PubSub with the Google SDK or gcloud. But if you are sending the message locally with a direct HTTP request, you need to encode the string in base64 beforehand

Eg. for "SELECT vendor_id FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY vendor_id;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type"

```
data: "U0VMRUNUIHZlbmRvcl9pZCBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgdmVuZG9yX2lkO1NFTEVDVCBwYXltZW50X3R5cGUsIFNVTSh0b3RhbF9hbW91bnQpIGFzIGNvbDEgRlJPTSBhY2hpbGlvLWRldi5ueWNfdHJpcHMudGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsIEdST1VQIEJZIHBheW1lbnRfdHlwZQ=="
```
