package main

import "net/http"

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir(app.staticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", app.HomeHandler)

	mux.HandleFunc("/quint/create", app.QuintCreateHandler)
	mux.HandleFunc("/quint/update", app.QuintUpdateHandler)
	mux.HandleFunc("/quint/delete", app.QuintDeleteHandler)
	mux.HandleFunc("/quint/list", app.QuintListHandler)
	mux.HandleFunc("/quint/get", app.QuintGetHandler)

	return secureHeaders(mux)
}
