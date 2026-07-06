package main

import (
	"net/http"
)

func (app *application) render(w http.ResponseWriter, fileName string, data interface{}) {
	if app.tp == nil {
		http.Error(w, "Template renderer not initialized", http.StatusInternalServerError)
		return
	}
	app.tp.Render(w, fileName, data)

}
