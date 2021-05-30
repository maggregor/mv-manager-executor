
gcloud pubsub topics publish mvExecutorTopic --message "Runner" --attribute="cmdType"="workspace","projectId"="myproject"
gcloud pubsub topics publish mvExecutorTopic --message "Runner" --attribute="cmdType"="apply","projectId"="myproject","regionId"="myregion","datasetId"="mydataset"