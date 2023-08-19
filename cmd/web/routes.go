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

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.HomeHandler))
	router.Handler(http.MethodGet, "/quint/view/:id", dynamic.ThenFunc(app.QuintViewHandler))
	router.Handler(http.MethodGet, "/quint/create", dynamic.ThenFunc(app.QuintCreateHandler))
	router.Handler(http.MethodPost, "/quint/create", dynamic.ThenFunc(app.QuintCreatePostHandler))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
