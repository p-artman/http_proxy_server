package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Front-end: https://app.uizard.io/prototypes/veAr86OJ60CxPrO51Jg1

func main() {
	registerRoutes()
	tasks = make(taskIDs)
	log.Printf("starting the server on %v:%v...", servHost, servPort)

	var started bool

	// Goroutine to log the server startup success
	go func() {
		time.Sleep(1 * time.Second) // Wait for a second as a temp solution
		if started {
			log.Printf("the server is successfully started.")
		}
	}()

	started = true // temp implementation: ListenAndServe blocks execution
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", servHost, servPort), nil)

	if err != nil {
		started = false // Signal server startup failure
		log.Fatalf("failed to start the server: %v", err)
	}
}
