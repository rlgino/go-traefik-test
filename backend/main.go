package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

var id string

func main() {
	fmt.Println("Hola mundo!")
	id = uuid.New().String()

	http.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/hello", HandleHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello "+id+"!")
}
