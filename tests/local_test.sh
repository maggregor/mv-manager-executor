# data is: SELECT payment_type, SUM(total_amount) as col1 FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY payment_type
echo '{
    "message": {
        "attributes": {
            "accessToken": "ya29.a0AfH6SMBT8SvdIL0dPSWPu_-WrLJ3l1eKIMT_PxocbY-WM11qQ_N_duU93dlfdi6MTMFMki8pLHfr7gFxak0ay8bh1zgKjnlNSormllvhjj7Z86dLsl9_zOydAh17jRLNM4k7BhoSw9kDOfQbpjpaXcsuT9yn_Q",
            "cmdType": "apply",
            "projectId": "achilio-dev",
            "regionId": "europe-west-1",
            "datasetId": "nyc_trips"
        },
        "data": "U0VMRUNUIHBheW1lbnRfdHlwZSwgU1VNKHRvdGFsX2Ftb3VudCkgYXMgY29sMSBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgcGF5bWVudF90eXBl",
        "messageId": "2070443601311540",
        "message_id": "2070443601311540",
        "publishTime": "2021-02-26T19:13:55.749Z",
        "publish_time": "2021-02-26T19:13:55.749Z"
    },
   "subscription": "projects/myproject/subscriptions/mysubscription"
}' | http POST localhost:8080