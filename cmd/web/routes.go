package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileserver := http.FileServer(http.Dir(app.staticDir))
	router.Handler(http.MethodGet, "/static/", http.StripPrefix("/static", fileserver))

	router.HandlerFunc(http.MethodGet, "/", app.HomeHandler)
	router.HandlerFunc(http.MethodGet, "/quint/view/:id", app.QuintViewHandler)
	router.HandlerFunc(http.MethodGet, "/quint/create", app.QuintCreateHandler)
	router.HandlerFunc(http.MethodPost, "/quint/create", app.QuintCreatePostHandler)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
