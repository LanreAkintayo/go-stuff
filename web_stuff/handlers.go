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

func home(w http.ResponseWriter, r *http.Request) {
	homeContent := fmt.Sprintf(htmlContent, "Home", "<h1>this is the home page</h1>")
	w.Write([]byte(homeContent))
}

func about(w http.ResponseWriter, r *http.Request) {
	aboutContent := fmt.Sprintf(htmlContent, "About", "<h1>this is the about page</h1>")
	w.Write([]byte(aboutContent))
}

func contact(w http.ResponseWriter, r *http.Request) {
	contactContent := fmt.Sprintf(htmlContent, "Contact", "<h1>this is the contact page</h1>")
	w.Write([]byte(contactContent))
}
