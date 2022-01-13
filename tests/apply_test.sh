# data is: SELECT vendor_id FROM achilio-dev.nyc_trips.tlc_yellow_trips_2015_small GROUP BY vendor_id", "SELECT (dropoff_latitude) IS NULL AS j17a,NOT ((pickup_latitude) IS NULL) AS a1894145264,pickup_latitude >= (NUMERIC \"38\") AS x92f,pickup_latitude <= (NUMERIC \"45\") AS c801,dropoff_longitude >= (NUMERIC \"-80\") AS p4ac,dropoff_longitude AS t829,NOT ((dropoff_latitude) IS NULL) AS l2ed,passenger_count AS bd99,NOT ((dropoff_longitude) IS NULL) AS kee0,(dropoff_longitude) IS NULL AS e01f,pickup_longitude AS vc89,dropoff_latitude AS c21c,dropoff_longitude <= (NUMERIC \"-70\") AS r1c6,pickup_latitude AS sb30,(NOT ((dropoff_latitude) IS NULL)) AND (NOT ((dropoff_longitude) IS NULL)) AND (NOT ((pickup_latitude) IS NULL)) AND (NOT ((pickup_longitude) IS NULL)) AND ((dropoff_longitude <= (NUMERIC \"-70\")) AND (dropoff_longitude >= (NUMERIC \"-80\")) AND (pickup_latitude >= (NUMERIC \"38\")) AND (pickup_latitude <= (NUMERIC \"45\"))) AS fb4d,(pickup_latitude) IS NULL AS o5cf,(pickup_longitude) IS NULL AS u75f,NOT ((pickup_longitude) IS NULL) AS w6fc,(dropoff_longitude <= (NUMERIC \"-70\")) AND (dropoff_longitude >= (NUMERIC \"-80\")) AND (dropoff_longitude >= (NUMERIC \"38\")) AND (dropoff_longitude <= (NUMERIC \"45\")) AS tc81 ,  FROM `nyc_trips`.`tlc_yellow_trips_2015_small` GROUP BY j17a,a1894145264,x92f,c801,p4ac,t829,l2ed,bd99,kee0,e01f,vc89,c21c,r1c6,sb30,fb4d,o5cf,u75f,w6fc,tc81
echo '{
    "message": {
        "attributes": {
            "accessToken": "accessTokenValue",
            "cmdType": "apply",
            "projectId": "achilio-dev",
            "datasetName": "nyc_trips",
        },
        "data": "U0VMRUNUIHZlbmRvcl9pZCBGUk9NIGFjaGlsaW8tZGV2Lm55Y190cmlwcy50bGNfeWVsbG93X3RyaXBzXzIwMTVfc21hbGwgR1JPVVAgQlkgdmVuZG9yX2lkIiwgIlNFTEVDVCAoZHJvcG9mZl9sYXRpdHVkZSkgSVMgTlVMTCBBUyBqMTdhLE5PVCAoKHBpY2t1cF9sYXRpdHVkZSkgSVMgTlVMTCkgQVMgYTE4OTQxNDUyNjQscGlja3VwX2xhdGl0dWRlID49IChOVU1FUklDIFwiMzhcIikgQVMgeDkyZixwaWNrdXBfbGF0aXR1ZGUgPD0gKE5VTUVSSUMgXCI0NVwiKSBBUyBjODAxLGRyb3BvZmZfbG9uZ2l0dWRlID49IChOVU1FUklDIFwiLTgwXCIpIEFTIHA0YWMsZHJvcG9mZl9sb25naXR1ZGUgQVMgdDgyOSxOT1QgKChkcm9wb2ZmX2xhdGl0dWRlKSBJUyBOVUxMKSBBUyBsMmVkLHBhc3Nlbmdlcl9jb3VudCBBUyBiZDk5LE5PVCAoKGRyb3BvZmZfbG9uZ2l0dWRlKSBJUyBOVUxMKSBBUyBrZWUwLChkcm9wb2ZmX2xvbmdpdHVkZSkgSVMgTlVMTCBBUyBlMDFmLHBpY2t1cF9sb25naXR1ZGUgQVMgdmM4OSxkcm9wb2ZmX2xhdGl0dWRlIEFTIGMyMWMsZHJvcG9mZl9sb25naXR1ZGUgPD0gKE5VTUVSSUMgXCItNzBcIikgQVMgcjFjNixwaWNrdXBfbGF0aXR1ZGUgQVMgc2IzMCwoTk9UICgoZHJvcG9mZl9sYXRpdHVkZSkgSVMgTlVMTCkpIEFORCAoTk9UICgoZHJvcG9mZl9sb25naXR1ZGUpIElTIE5VTEwpKSBBTkQgKE5PVCAoKHBpY2t1cF9sYXRpdHVkZSkgSVMgTlVMTCkpIEFORCAoTk9UICgocGlja3VwX2xvbmdpdHVkZSkgSVMgTlVMTCkpIEFORCAoKGRyb3BvZmZfbG9uZ2l0dWRlIDw9IChOVU1FUklDIFwiLTcwXCIpKSBBTkQgKGRyb3BvZmZfbG9uZ2l0dWRlID49IChOVU1FUklDIFwiLTgwXCIpKSBBTkQgKHBpY2t1cF9sYXRpdHVkZSA+PSAoTlVNRVJJQyBcIjM4XCIpKSBBTkQgKHBpY2t1cF9sYXRpdHVkZSA8PSAoTlVNRVJJQyBcIjQ1XCIpKSkgQVMgZmI0ZCwocGlja3VwX2xhdGl0dWRlKSBJUyBOVUxMIEFTIG81Y2YsKHBpY2t1cF9sb25naXR1ZGUpIElTIE5VTEwgQVMgdTc1ZixOT1QgKChwaWNrdXBfbG9uZ2l0dWRlKSBJUyBOVUxMKSBBUyB3NmZjLChkcm9wb2ZmX2xvbmdpdHVkZSA8PSAoTlVNRVJJQyBcIi03MFwiKSkgQU5EIChkcm9wb2ZmX2xvbmdpdHVkZSA+PSAoTlVNRVJJQyBcIi04MFwiKSkgQU5EIChkcm9wb2ZmX2xvbmdpdHVkZSA+PSAoTlVNRVJJQyBcIjM4XCIpKSBBTkQgKGRyb3BvZmZfbG9uZ2l0dWRlIDw9IChOVU1FUklDIFwiNDVcIikpIEFTIHRjODEgLCAgRlJPTSBgbnljX3RyaXBzYC5gdGxjX3llbGxvd190cmlwc18yMDE1X3NtYWxsYCBHUk9VUCBCWSBqMTdhLGExODk0MTQ1MjY0LHg5MmYsYzgwMSxwNGFjLHQ4MjksbDJlZCxiZDk5LGtlZTAsZTAxZix2Yzg5LGMyMWMscjFjNixzYjMwLGZiNGQsbzVjZix1NzVmLHc2ZmMsdGM4MQ==",
        "messageId": "2070443601311540",
        "message_id": "2070443601311540",
        "publishTime": "2021-02-26T19:13:55.749Z",
        "publish_time": "2021-02-26T19:13:55.749Z"
    },
   "subscription": "projects/myproject/subscriptions/mysubscription"
}' | http POST localhost:8080