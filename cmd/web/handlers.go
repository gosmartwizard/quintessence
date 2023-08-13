package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gosmartwizard/quintessence/internal/models"
)

func (app *application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	quints, err := app.quints.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, quint := range quints {
		fmt.Fprintf(w, "%+v\n", quint)
	}

	/* ts, err := template.ParseFiles("./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl")
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	} */
}

func (app *application) QuintCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	title := "GO Interfaces"
	content := "An interface in Go is a type defined using a set of method signatures. The interface defines the behavior of a similar type of object"
	expires := 7

	id, err := app.quints.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/quint/get?id=%d", id), http.StatusSeeOther)
}

func (app *application) QuintGetHandler(w http.ResponseWriter, r *http.Request) {

	/* 	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, "Method not supported", http.StatusNotFound)
		return
	} */

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

	quint, err := app.quints.Get(qid)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", quint)
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
	quints, err := app.quints.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, quint := range quints {
		fmt.Fprintf(w, "%+v\n", quint)
	}
}

func (app *application) QuintUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", http.MethodPut)
		app.clientError(w, "Method not supported", http.StatusNotFound)
		return
	}
	w.Write([]byte("Successfully updated a quint"))
}
