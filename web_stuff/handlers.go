package main

// import "fmt"
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
	// app.infoLog.Printf("Session data: %v", app.session.Get(r, "userID"))
	app.render(w, "index.html", nil)
}
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	// app.session.Put(r, "userID", "Lanre")
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)
		form.Required("email", "password").MaxLength("email", 100).MaxLength("password", 100).MinLength("password", 6).Matches("email", EmailRX)

		email := r.FormValue("email")
		password := r.FormValue("password")

		app.infoLog.Printf("Email: %s | Password: %s", email, password)

	}
	app.render(w, "login.html", nil)
}
func (app *application) register(w http.ResponseWriter, r *http.Request) {
	app.render(w, "register.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, "about.html", nil)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, "contact.html", nil)
}
func (app *application) submit(w http.ResponseWriter, r *http.Request) {
	app.render(w, "submit.html", nil)

}
