// Sample run-pubsub is a Cloud Run service which handles Pub/Sub messages.
package main

import (
	"log"
	"net/http"
	"os"

	e "github.com/achilio/mv-manager-executor/executor"
)

func main() {
	http.HandleFunc("/", e.PubSub)
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	// Start HTTP server.
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
