# Achilio mv-manager-executor

[![codecov](https://codecov.io/gh/achilio/mv-manager-executor/branch/master/graph/badge.svg?token=OLM9U79QD4)](https://codecov.io/gh/achilio/mv-manager-executor)

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

The message data is a list of map with the following structure:

{
"datasetName": "queryStatement"
}

It is base64 encoded

Eg.
[{"mydataset1": "SELECT 1"},{"mydataset2": "SELECT 1"}]

```
data: "W3sibXlkYXRhc2V0MSI6ICJTRUxFQ1QgMSJ9LHsibXlkYXRhc2V0MiI6ICJTRUxFQ1QgMSJ9XQ=="
```
