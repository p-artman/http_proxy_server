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
	log.Printf("starting the server on :%d...", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil)
	log.Fatal(http.ListenAndServe(fmt.Sprint(serverPort), nil))
	// TODO: add check whether the server is actually started
	log.Printf("the server is successfully started.")
}
