package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	// initial version: hardcoded IPv4 address and port
	serverHost string = "127.0.0.1"
	serverPort uint16 = 4000
)

var (
	tasks taskIDs // map of all tasks: map[string]task
)

type taskIDs map[string]task

type task struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Body   []byte `json:"body,omitempty"`
}

// todo: implement an anonymous proxy HTTP-server logic:
/*
	• JSON-request parser
	• Client task manager: generates an ID and operate it
	• HTTP-request sender
	• Task status checker
	•
*/

func newTask(ts taskIDs, r *http.Request) task {
	t := task{Status: "new"}
	// TODO: fix the following logic
	fmt.Fscan(rand.Reader, &t.ID)
	t.ID = base64.URLEncoding.EncodeToString([]byte(t.ID))
	ts[t.ID] = t
	return t
}

func (t *task) marshalID() ([]byte, error) {
	var tmp struct {
		ID string `json:"id"`
	}
	tmp.ID = t.ID
	return json.Marshal(tmp)
}

// func (t *task) marshalBody() ([]byte, error) {
// 	// var tmp struct {
// 	// 	ID string `json:"id"`
// 	// }
// 	// tmp.ID = t.ID
// 	var tmp interface{map[string]interface{}}
// 	json.Unmarshal(t.Body, tmp)
// 	return json.Marshal(tmp)
// }

func (t *task) execute() {
	// TODO: add a shared channel and a goroutine for each task
	// TODO: implement a client side logic there with a buffer
	t.Status = "in progress"
	// go t.ID
	// tw, tr := http.ResponseWriter, *http.Request
}

func (t *task) getStatus(w http.ResponseWriter) {
	// TODO: check and generate a JSON status response
	// TODO: add a new method for t to marshal with response
	sts, err := json.Marshal(t.Status)
	if err != nil {
		panic(err) // TODO: remove panic
	}
	fmt.Fprint(w, sts)
}

func validBodyJSON(w http.ResponseWriter, r *http.Request) bool {
	byteData, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if !json.Valid(byteData) {
		errMsg := fmt.Sprintf("request body doesn't contain a valid JSON.\n%v", string(byteData))
		log.Printf("[ERRR] http 400 error - %s", errMsg)
		// log.Printf("[DEBG] Request body contents: \n%v", string(byteData))
		http.Error(w, "Error: "+errMsg, http.StatusBadRequest)
		return false
	}
	log.Printf("[INFO] http 200 code - request body received, contents: \n%v", string(byteData))
	return true
}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/task":
// 		taskHandlerFunc(w, r)
// 		return
// 	default:
// 		http.NotFound(w, r)
// 	}
// }

// func executeTemplate(w http.ResponseWriter, filepath string) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	tpl, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		log.Printf("processing template: %v", err)
// 		http.Error(w, "An error happened processing the template.", http.StatusInternalServerError)
// 		return
// 	}
// 	err = tpl.Execute(w, nil)
// 	if err != nil {
// 		log.Printf("executing template: %v", err)
// 		http.Error(w, "An error happened executing the template.", http.StatusInternalServerError)
// 		return
// 	}
// }
