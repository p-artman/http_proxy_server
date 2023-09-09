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
	log.Printf("Starting the server on :%d...", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil)
	log.Fatal(http.ListenAndServe(fmt.Sprint(serverPort), nil))
}
