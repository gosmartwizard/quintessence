package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, msg string, status int) {
	http.Error(w, msg, status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, "Not Found", http.StatusNotFound)
}
