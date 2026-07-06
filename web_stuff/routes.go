package main

import (
	"net/http"
)


func (app *application) routes() http.Handler{
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(app.publicPath))
	mux.Handle("/public/", http.StripPrefix("/public/", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/contact", app.contact)
	mux.HandleFunc("/login", app.login)
	mux.HandleFunc("/register", app.register)

	handler := app.recover(app.logger(app.session.Enable(mux)))
	
	return handler
}
