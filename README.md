# Achilio mv-manager-executor

## Roadmap

- [x] Handling notification message to build the command to execute
- [ ] Add terraform module to actually execute

## Run locally

```
docker build -t local/mv-manager-executor .
docker run local/mv-manager-executor -p 8080
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

- "cmdType": enum("apply", "workspace)
- "projectId": string
- "regionId": string // Not required if cmdType is workspace
- "datasetId": string // Not required if cmdType is workspace
- "accessToken: string // Not required for testing purpose, but will be when in production. If not provided, you need to have a way to pass credentials to the executor with a json service account file.