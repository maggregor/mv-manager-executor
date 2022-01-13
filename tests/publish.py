from google.cloud import pubsub_v1

# TODO(developer)
project_id = "achilio-dev"
topic_id = "mvExecutorTopic"

publisher = pubsub_v1.PublisherClient()
topic_path = publisher.topic_path(project_id, topic_id)

data = "Message number 1"
# Data must be a bytestring
data = data.encode("utf-8")
# Add two attributes, origin and username, to the message
future = publisher.publish(
    topic_path,
    data,
    cmdType="apply",
    projectId="achilio-dev",
    datasetName="nyc_trips",
    queries=["SELECT 1", "SELECT 2"],
)
print(future.result())

print(f"Published messages with custom attributes to {topic_path}.")