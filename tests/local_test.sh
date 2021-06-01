echo '{
    "message": {
        "attributes": {
            "accessToken": "accessTokenValue",
            "cmdType": "apply",
            "projectId": "myprojectid",
            "regionId": "myregion",
            "datasetId": "mydataset"
        },
        "data": "SELECT vendor_id FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY vendor_id",
        "messageId": "2070443601311540",
        "message_id": "2070443601311540",
        "publishTime": "2021-02-26T19:13:55.749Z",
        "publish_time": "2021-02-26T19:13:55.749Z"
    },
   "subscription": "projects/myproject/subscriptions/mysubscription"
}' | http POST localhost:8080