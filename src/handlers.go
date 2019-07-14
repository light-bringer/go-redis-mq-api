package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	mux "github.com/julienschmidt/httprouter"
)

// Logger logs the method, URI, header, and the dispatch time of the request.
func Logger(r *http.Request) {
	start := time.Now()
	log.Printf(
		"%s\t%s\t%q\t%s",
		r.Method,
		r.RequestURI,
		r.Header,
		time.Since(start),
	)
}

// Index handler handles the index at "/" and writes a welcome message.
func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to user creation service</h1>")
}


func UserCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	Logger(r)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Save JSON to Post struct (should this be a pointer?)
	var user UserObject
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}

	CreateUser(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
