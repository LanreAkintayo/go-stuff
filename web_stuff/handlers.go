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
	app.infoLog.Printf("%s %s", r.Method, r.URL.Path)
	homeContent := fmt.Sprintf(htmlContent, "Home", "<h1>this is the home page</h1>")
	w.Write([]byte(homeContent))
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
