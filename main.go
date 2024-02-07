package main

import (
	"fmt"
	"log"
	"net/http"
)

// Front-end: https://app.uizard.io/prototypes/veAr86OJ60CxPrO51Jg1

func main() {
	registerRoutes()
	tasks = make(taskIDs)
	log.Printf("starting the server on %v:%v...", serverHost, serverPort)
	http.ListenAndServe(fmt.Sprintf("%v:%v", serverHost, serverPort), nil)
	log.Fatal(http.ListenAndServe(fmt.Sprint(serverPort), nil))
	// TODO: add a verification step to ensure that the server is actually started (the line above is never initiated for now)
	log.Printf("the server is successfully started.")
}
