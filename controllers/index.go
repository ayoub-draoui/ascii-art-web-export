package controllers

import (
	"html/template"
	"net/http"

	"functions/functions"
)

var tmpl *template.Template

// output string

func Index(w http.ResponseWriter, r *http.Request) {
	d := functions.Data{}
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	// ERROR 404 Page Not Found
	if r.URL.Path != "/" {
		d.ErrNum = http.StatusNotFound
		d.ErrTxt = "Page Not Found"
		functions.HandlerError(w, r, &d)
		return
	}
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/ascii", http.StatusMovedPermanently)
	} else {
		// 400 Bad Request
		d.ErrNum = http.StatusBadRequest
		d.ErrTxt = "Bad Request"
		functions.HandlerError(w, r, &d)
		return
	}
}
