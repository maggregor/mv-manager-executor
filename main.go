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
	port := os.Getenv("ACHILIO_EXECUTOR_PORT")
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
