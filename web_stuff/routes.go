package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	defaultMiddleware := alice.New(app.recover, app.logger)
	secureMiddleware := alice.New(app.session.Enable)

	fileServer := http.FileServer(http.Dir(app.publicPath))
	mux.Handle("/public/", http.StripPrefix("/public/", fileServer))
	mux.Handle("/", secureMiddleware.ThenFunc(app.home))
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/contact", app.contact)
	mux.Handle("/login", secureMiddleware.ThenFunc(app.login))
	mux.Handle("/register", secureMiddleware.ThenFunc(app.register))
	mux.Handle("/submit", secureMiddleware.Append(app.requireAuth).ThenFunc(app.submit))

	handler := defaultMiddleware.Then(mux)

	return handler
}
