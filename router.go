package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func registerRoutes() {
	http.HandleFunc("/", router)
	http.HandleFunc("/task", router)
	http.HandleFunc("/task/", taskHandler)
}

func router(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] route handler, parse URL...")
	log.Printf("[DEBUG] request endpoint: %#v\n", r.URL.Path[1:])
	switch r.URL.Path {
	case "/task":
		taskHandler(w, r)
		return
	case "/favicon.ico":
		faviconHandler(w, r)
		return
	case "/":
		homeHandler(w, r)
		return
	default:
		errMsg := fmt.Sprint("wrong URL: the requested path is not present.")
		log.Printf("[ERROR] %v\n", errMsg)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: add "greetings" template, maybe even a frontend
	// tplPath := filepath.Join("templates", "home.gohtml")
	// executeTemplate(w, tplPath)
	log.Printf("[INFO] home handler, parse method %v...\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		log.Printf("[INFO] sending home page template response...")
		fmt.Fprintln(w, "<h1>Server is up and running</h1>")
	default:
		http.Error(w, http.StatusText(418), http.StatusTeapot)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	// NOTE: icon source: https://favicon.io/favicon-generator/
	w.Header().Set("Content-Type", "image/x-icon")
	log.Printf("[INFO] favicon handler, parse method %v...\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "templates/favicon.ico")
	default:
		http.Error(w, http.StatusText(418), http.StatusTeapot)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pathParams bool
	switch endpoint := r.URL.Path[1:]; endpoint {
	case "task":
		pathParams = false
	default:
		pathParams = true
	}
	log.Println("[TRACE] task handler, flag: ", pathParams)
	log.Println("[TRACE] task handler, method: ", r.Method)
	log.Println("[TRACE] task handler, URL: ", r.URL)
	switch r.Method {
	case http.MethodPost:
		if pathParams {
			errMsg := fmt.Sprint("wrong URL: only \"/task\" is allowed for POST.")
			log.Printf("[ERROR] %v\n", errMsg)
			http.Error(w, errMsg, http.StatusBadRequest)
			return
		}
		log.Println("[INFO] task handler, post method")
		if !validBodyJSON(w, r) {
			return
		}
		t := newTask(tasks, r)
		log.Printf("[DEBUG] task: %#v", t)
		// TODO: add a JSON formatted taskID response
		out, err := t.marshalID()
		if err != nil {
			panic(err) // TODO: remove panic
		}
		log.Printf("[DEBUG] task in json: %s", out)
		fmt.Fprintf(w, "%s", out)
		t.execute()
	case http.MethodGet:
		log.Println("[INFO] task handler, get method, total:", len(tasks))
		if len(tasks) != 0 {
			tsk := strings.TrimPrefix(r.URL.Path, "/task/")
			log.Printf("[DEBUG] taskID: %v\n", tsk)
			if _, ok := tasks[tsk]; !ok {
				errMsg := fmt.Sprint("Error: the requested task doesn't exist!")
				http.Error(w, errMsg, http.StatusBadRequest)
			}
			t := tasks[tsk]
			fmt.Fprint(w, t.Status)
		} else {
			fmt.Fprintf(w, "Error: no tasks are present!")
		}
	default:
		http.Error(w, http.StatusText(418), http.StatusTeapot)
	}
}
