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

const (
	loggedInUserKey = "logged_in_user"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// app.infoLog.Printf("Session data: %v", app.session.Get(r, "userID"))
	app.render(w, r, "index.html", nil)
}
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	// app.session.Put(r, "userID", "Lanre")
	app.infoLog.Printf("Session key: %s", app.session.GetString(r, loggedInUserKey))
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)
		form.Required("email", "password").
			MaxLength("email", 100).
			MaxLength("password", 100).
			MinLength("password", 6).
			MinLength("email", 3).
			Matches("email", EmailRX)

		if !form.Valid() {
			form.Errors.Add("generic", "The data you submitted is not valid")
			app.errorLog.Printf("Invalid form: %+v", form.Errors)
			app.render(w, r, "login.html", &templateData{Form: form})
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		userID, err := app.userRepo.Authenticate(email, password)
		if err != nil {
			form.Errors.Add("generic", "Invalid credentials")
			app.render(w, r, "login.html", &templateData{Form: form})
			return
		}

		// We put the userID in session.
		app.session.Put(r, loggedInUserKey, userID)

		app.infoLog.Printf("User %d logged in", userID)

		redirectURL := r.FormValue("redirectTo")
		if redirectURL == "" {
			redirectURL = "/submit"
		}

		http.Redirect(w,r, redirectURL, http.StatusSeeOther)


	}
	app.render(w, r, "login.html", &templateData{Form: NewForm(r.PostForm)})
}
func (app *application) register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		form := NewForm(r.PostForm)
		form.Required("name", "email", "password").
			MaxLength("name", 100).
			MaxLength("email", 100).
			MinLength("password", 6).
			MinLength("email", 3).
			Matches("email", EmailRX)
		if !form.Valid() {
			form.Errors.Add("generic", "The data you submitted is not valid")
			app.errorLog.Printf("Invalid form: %+v", form.Errors)
			app.render(w, r, "register.html", &templateData{Form: form})
			return
		}
		
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		avatar := r.FormValue("avatar")
		
		userID, err := app.userRepo.CreateUser(name, email, password, avatar)
		if err != nil {
			app.errorLog.Printf("Error creating user: %v", err.Error())
			form.Errors.Add("generic", "Failed to create account")
			app.render(w, r, "register.html", &templateData{Form: form})
			return
		}
		
		// app.session.Put(r, loggedInUserKey, userID)
		
		app.infoLog.Printf("User %d registered", userID)
		
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	app.render(w, r, "register.html", &templateData{Form: NewForm(r.PostForm)})
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.html", nil)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.html", nil)
}
func (app *application) submit(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "submit.html", nil)

}
