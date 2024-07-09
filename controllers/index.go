package controllers

import (
	"html/template"
	"net/http"

	"functions/functions"
)

var tmpl *template.Template

// output string

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	// ERROR 404 Page Not Found
	if r.URL.Path != "/" {
		functions.MessageError(w, r, http.StatusNotFound, "Page Not Found") // Handle the error and return
		return
	}
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/ascii", http.StatusMovedPermanently)
	} else {
		// 400 Bad Request
		functions.MessageError(w, r, http.StatusBadRequest, "Bad Request") // Handle the error and return
		return
	}
}
