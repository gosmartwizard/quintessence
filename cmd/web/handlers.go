package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gosmartwizard/quintessence/internal/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) HomeHandler(w http.ResponseWriter, r *http.Request) {

	quints, err := app.quints.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Quints = quints

	app.render(w, http.StatusOK, "home.tmpl", data)
}

func (app *application) QuintViewHandler(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	sid := params.ByName("id")
	if sid == "" {
		app.clientError(w, "There is no id in the query", http.StatusBadRequest)
		return
	}
	qid, err := strconv.Atoi(sid)
	if err != nil {
		msg := fmt.Sprintf("Failed to convert ID : %v ", sid)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	} else if qid < 1 {
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

	data := app.newTemplateData(r)
	data.Quint = quint

	app.render(w, http.StatusOK, "view.tmpl", data)
}

func (app *application) QuintCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the form for creating a new Qunit..."))
}

func (app *application) QuintCreatePostHandler(w http.ResponseWriter, r *http.Request) {

	title := "Sadhguru"
	content := "Live life in a blissful way"
	expires := 7

	id, err := app.quints.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/quint/view/%d", id), http.StatusSeeOther)
}
