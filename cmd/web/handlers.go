package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	ts, err := template.ParseFiles("./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl")
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) QuintCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Successfully created a new quint"))
}

func (app *application) QuintGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, "Method not supported", http.StatusNotFound)
		return
	}

	sid := r.URL.Query().Get("id")
	if sid == "" {
		app.clientError(w, "There is no id in the query", http.StatusBadRequest)
		return
	}
	qid, err := strconv.Atoi(sid)
	if err != nil {
		msg := fmt.Sprintf("Failed to convert ID : %v ", sid)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	} else if qid < 0 {
		msg := fmt.Sprintf("Not a valid ID : %d ", qid)
		app.clientError(w, msg, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Successfully retrieved a quint %d", qid)
}

func (app *application) QuintDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		app.clientError(w, "Method not supported", http.StatusNotFound)
		return
	}
	w.Write([]byte("Successfully deleted a quint"))
}

func (app *application) QuintListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, "Method not supported", http.StatusNotFound)
		return
	}
	w.Write([]byte("Successfully list a quint"))
}

func (app *application) QuintUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodPut)
		app.clientError(w, "Method not supported", http.StatusNotFound)
		return
	}
	w.Write([]byte("Successfully updated a quint"))
}
