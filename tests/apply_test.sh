echo '{
    "message": {
        "attributes": {
            "accessToken": "accessTokenValue",
            "cmdType": "apply",
            "projectId": "achilio-dev",
            "regionId": "europe-west-1",
            "datasetId": "nyc_trips",
            "queries": ["SELECT vendor_id FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY vendor_id", "SELECT (dropoff_latitude) IS NULL AS j17a,NOT ((pickup_latitude) IS NULL) AS a1894145264,pickup_latitude >= (NUMERIC \"38\") AS x92f,pickup_latitude <= (NUMERIC \"45\") AS c801,dropoff_longitude >= (NUMERIC \"-80\") AS p4ac,dropoff_longitude AS t829,NOT ((dropoff_latitude) IS NULL) AS l2ed,passenger_count AS bd99,NOT ((dropoff_longitude) IS NULL) AS kee0,(dropoff_longitude) IS NULL AS e01f,pickup_longitude AS vc89,dropoff_latitude AS c21c,dropoff_longitude <= (NUMERIC \"-70\") AS r1c6,pickup_latitude AS sb30,(NOT ((dropoff_latitude) IS NULL)) AND (NOT ((dropoff_longitude) IS NULL)) AND (NOT ((pickup_latitude) IS NULL)) AND (NOT ((pickup_longitude) IS NULL)) AND ((dropoff_longitude <= (NUMERIC \"-70\")) AND (dropoff_longitude >= (NUMERIC \"-80\")) AND (pickup_latitude >= (NUMERIC \"38\")) AND (pickup_latitude <= (NUMERIC \"45\"))) AS fb4d,(pickup_latitude) IS NULL AS o5cf,(pickup_longitude) IS NULL AS u75f,NOT ((pickup_longitude) IS NULL) AS w6fc,(dropoff_longitude <= (NUMERIC \"-70\")) AND (dropoff_longitude >= (NUMERIC \"-80\")) AND (dropoff_longitude >= (NUMERIC \"38\")) AND (dropoff_longitude <= (NUMERIC \"45\")) AS tc81 ,  FROM `nyc_trips`.`tlc_yellow_trips_2015_small` GROUP BY j17a,a1894145264,x92f,c801,p4ac,t829,l2ed,bd99,kee0,e01f,vc89,c21c,r1c6,sb30,fb4d,o5cf,u75f,w6fc,tc81"]
        },
        "data": "SGVsbG8gQ2xvdWQgUHViL1N1YiEgSGVyZSBpcyBteSBtZXNzYWdlIQ==",
        "messageId": "2070443601311540",
        "message_id": "2070443601311540",
        "publishTime": "2021-02-26T19:13:55.749Z",
        "publish_time": "2021-02-26T19:13:55.749Z"
    },
   "subscription": "projects/myproject/subscriptions/mysubscription"
}' | http POST localhost:8080