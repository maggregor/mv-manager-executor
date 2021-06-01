# Achilio mv-manager-executor

## Roadmap

- [x] Handling notification message to build the command to execute
- [x] Add terraform module to actually execute

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

- "cmdType": enum("apply", "workspace")
- "projectId": string
- "regionId": string // Not required if cmdType is workspace
- "datasetId": string // Not required if cmdType is workspace
- "accessToken: string // Not required for testing purpose, but will be when in production. If not provided, you need to have a way to pass credentials to the executor with a json service account file.

The message data is a list of ; separated SELECT queries representing the views to create.

Eg. data: "SELECT vendor_id FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY vendor_id;SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type"

