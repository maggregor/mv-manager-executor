# data is: [{"executor_integration_test":"SELECT COUNT(*) AS a_410576920 FROM `achilio-test`.`executor_integration_test`.`tlc_yellow_trips_2015_small`"}]
echo '{
    "message": {
        "attributes": {
            "accessToken": "ya29.A0ARrdaM8aOcSqtc1OB92eCIjLJtDDQ_tMZ6hVGq3mDoNZwh9RyEJKlp_FAcgXGZ2QmySK-XCtywx-k8PUgNK2eHKzupkc1qGkyscr1l9Mgwa5NMvjPTY4qv_MXiUTmbojv98Xs3I4AN4PvjZk04BBSM1861PJKA",
            "cmdType": "apply",
            "projectId": "achilio-test"
        },
        "data": "W3siZXhlY3V0b3JfaW50ZWdyYXRpb25fdGVzdCI6IlNFTEVDVCBDT1VOVCgqKSBBUyBhXzQxMDU3NjkyMCBGUk9NIGBhY2hpbGlvLXRlc3RgLmBleGVjdXRvcl9pbnRlZ3JhdGlvbl90ZXN0YC5gYV90YWJsZWAifV0=",
        "messageId": "2070443601311540",
        "message_id": "2070443601311540",
        "publishTime": "2021-02-26T19:13:55.749Z",
        "publish_time": "2021-02-26T19:13:55.749Z"
    },
   "subscription": "projects/myproject/subscriptions/mysubscription"
}' | http POST localhost:8080