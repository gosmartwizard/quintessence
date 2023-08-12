package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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

	sid := r.URL.Query().Get("id")
	if sid == "" {
		http.Error(w, "Empty Id", http.StatusBadRequest)
		return
	}
	qid, err := strconv.Atoi(sid)
	if err != nil {
		msg := fmt.Sprintf("Failed to convert ID : %v ", sid)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	} else if qid < 0 {
		msg := fmt.Sprintf("Not a valid ID : %d ", qid)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Successfully retrieved a quint %d", qid)
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
