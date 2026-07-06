package main

import "fmt"
import "net/http"

var htmlContent = `
	<!DOCTYPE html>
	<html>
	<head><title>%s</title></head>
	<body>
	%s
	</body>
	</html>
`

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Session data: %v", app.session.Get(r, "userID"))
	app.render(w, "index.html", nil)
}
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	app.session.Put(r, "userID", "Lanre")
	app.render(w, "login.html", nil)
}
func (app *application) register(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s %s", r.Method, r.URL.Path)
	app.render(w, "register.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s %s", r.Method, r.URL.Path)
	aboutContent := fmt.Sprintf(htmlContent, "About", "<h1>this is the about page</h1>")
	w.Write([]byte(aboutContent))
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s %s", r.Method, r.URL.Path)
	contactContent := fmt.Sprintf(htmlContent, "Contact", "<h1>this is the contact page</h1>")
	w.Write([]byte(contactContent))
}
