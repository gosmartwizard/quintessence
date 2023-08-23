package main

import (
	"net/http"

	"github.com/gosmartwizard/quintessence/ui"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	//fileserver := http.FileServer(http.Dir(app.staticDir))
	//router.Handler(http.MethodGet, "/static/", http.StripPrefix("/static", fileserver))

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	router.Handler(http.MethodGet, "/about", dynamic.ThenFunc(app.AboutHandler))
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.HomeHandler))
	router.Handler(http.MethodGet, "/quint/view/:id", dynamic.ThenFunc(app.QuintViewHandler))
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))

	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/quint/create", protected.ThenFunc(app.QuintCreateHandler))
	router.Handler(http.MethodPost, "/quint/create", protected.ThenFunc(app.QuintCreatePostHandler))
	router.Handler(http.MethodGet, "/account/view", protected.ThenFunc(app.AccountHandler))
	router.Handler(http.MethodGet, "/account/password/update", protected.ThenFunc(app.accountPasswordUpdate))
	router.Handler(http.MethodPost, "/account/password/update", protected.ThenFunc(app.accountPasswordUpdatePost))
	router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
