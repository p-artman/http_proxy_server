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
	http.HandleFunc("/task/", taskHandler) // to ignore trailing slash
}

func router(w http.ResponseWriter, r *http.Request) {
	log.Println("*** http request received, processing...")
	log.Println("[DEBG] default http route handler is called, parse URL:")
	log.Printf("[INFO] method: %#v, endpoint is %#v, processing...\n", r.Method, r.URL.Path[1:])
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
		log.Printf("[ERRR] http 404 error - %v\n", errMsg)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: add "greetings" template, maybe even a frontend
	// tplPath := filepath.Join("templates", "home.gohtml")
	// executeTemplate(w, tplPath)
	log.Printf("[DEBG] home handler is called, processing \"%v\" method...\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		log.Printf("[INFO] http 200 code - sending home page template response.")
		fmt.Fprintln(w, "<h1>Server is up and running</h1>")
	default:
		log.Printf("[ERRR] http 418 error - method is not supported.")
		http.Error(w, http.StatusText(418), http.StatusTeapot)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	// NOTE: icon source: https://favicon.io/favicon-generator/
	w.Header().Set("Content-Type", "image/x-icon")
	log.Printf("[DEBG] favicon handler is called, processing \"%v\" method...\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "templates/favicon.ico")
		log.Printf("[INFO] http 200 code - favicon picture is returned.")
	default:
		log.Printf("[ERRR] http 418 error - method is not supported.")
		http.Error(w, http.StatusText(418), http.StatusTeapot)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pathArgs bool
	switch endpoint := r.URL.Path[1:]; endpoint {
	case "task":
		pathArgs = false
	default:
		pathArgs = true // when we request for /task/{{taskID}} for example
	}
	log.Println("*** http request received, processing...")
	log.Println("[DEBG] task handler is called, parse URL...")
	log.Printf("[INFO] method: %#v, endpoint is %#v, processing...\n", r.Method, r.URL.Path[1:])
	log.Println("[TRCE] URL path contain additional args:", pathArgs)
	switch r.Method {
	case http.MethodPost:
		if pathArgs {
			errMsg := fmt.Sprint("wrong URL: only \"/task\" endpoint is allowed for \"POST\" method.")
			log.Printf("[ERRR] http 400 error - %v\n", errMsg)
			http.Error(w, "Error: "+errMsg, http.StatusBadRequest)
			return
		}
		log.Println("[TRCE] handle \"POST\" method")
		if !validBodyJSON(w, r) {
			return
		}
		t := newTask(tasks, r)
		log.Printf("[DEBG] new task is added: \n%#v", t)
		// TODO: add a JSON formatted taskID response
		out, err := t.marshalID()
		if err != nil {
			panic(err) // TODO: remove panic
		}
		log.Printf("[DEBG] the task is converted into json: \n%s", out)
		// log.Printf("[DEBG] the task is converted into json: \n%s", json.Unmarshal(out, v)) // TODO: show formatted json in log
		fmt.Fprintf(w, "%s", out)
		t.execute()
	case http.MethodGet:
		log.Println("[TRCE] handle \"GET\" method, tasks total:", len(tasks))
		if len(tasks) != 0 {
			tsk := strings.TrimPrefix(r.URL.Path, "/task/")
			log.Printf("[DEBG] taskID: %v\n", tsk)
			if _, ok := tasks[tsk]; !ok {
				errMsg := fmt.Sprint("the requested task doesn't exist!")
				log.Printf("[ERRR] http 400 error - %s", errMsg)
				http.Error(w, "Error: "+errMsg, http.StatusBadRequest)
			}
			t := tasks[tsk]
			fmt.Fprint(w, t.Status)
		} else {
			// fmt.Fprintf(w, "Error: no tasks are present!")
			errMsg := fmt.Sprint("no tasks are present!")
			log.Printf("[ERRR] http 400 error - %s", errMsg)
			http.Error(w, "Error: "+errMsg, http.StatusBadRequest)
		}
	default:
		log.Printf("[ERRR] http 418 error - method is not supported.")
		http.Error(w, http.StatusText(418), http.StatusTeapot)
	}
}
