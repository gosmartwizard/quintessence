package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

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

	data := app.newTemplateData(r)

	data.Form = quintCreateForm{
		Expires: 365,
	}

	app.render(w, http.StatusOK, "create.tmpl", data)
}

type quintCreateForm struct {
	Title       string
	Content     string
	Expires     int
	FieldErrors map[string]string
}

func (app *application) QuintCreatePostHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, "Unable to Parse Form data", http.StatusBadRequest)
		return
	}

	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, "expires is not valid", http.StatusBadRequest)
		return
	}

	form := quintCreateForm{
		Title:       r.PostForm.Get("title"),
		Content:     r.PostForm.Get("content"),
		Expires:     expires,
		FieldErrors: map[string]string{},
	}

	if strings.TrimSpace(form.Title) == "" {
		form.FieldErrors["title"] = "Title field cannot be blank"
	} else if utf8.RuneCountInString(form.Title) > 100 {
		form.FieldErrors["title"] = "Title field cannot be more than 100 characters long"
	}
	if strings.TrimSpace(form.Content) == "" {
		form.FieldErrors["content"] = "Content field cannot be blank"
	}
	if expires != 1 && expires != 7 && expires != 365 {
		form.FieldErrors["expires"] = "Expires field must equal 1, 7 or 365"
	}

	if len(form.FieldErrors) > 0 {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}

	id, err := app.quints.Insert(form.Title, form.Content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/quint/view/%d", id), http.StatusSeeOther)
}
