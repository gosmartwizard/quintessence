package main

import (
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Welcome to the home page"))
}

func QuintCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Successfully created a new quint"))
}

func QuintGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Successfully retrieved a quint"))
}

func QuintDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Successfully deleted a quint"))
}

func QuintListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Successfully list a quint"))
}

func QuintUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodPut)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Successfully updated a quint"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)

	mux.HandleFunc("/quint/create", QuintCreateHandler)
	mux.HandleFunc("/quint/update", QuintUpdateHandler)
	mux.HandleFunc("/quint/delete", QuintDeleteHandler)
	mux.HandleFunc("/quint/list", QuintListHandler)
	mux.HandleFunc("/quint/get", QuintGetHandler)

	log.Print("Server listening on port : 4949 ")

	err := http.ListenAndServe(":4949", mux)
	log.Fatal(err)
}
