package main

import (
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)

	log.Print("Server listening on port : 4949 ")

	err := http.ListenAndServe(":4949", mux)
	log.Fatal(err)
}
