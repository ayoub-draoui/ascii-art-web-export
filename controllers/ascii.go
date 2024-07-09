package controllers

import (
	"html/template"
	"net/http"

	"functions/functions"
)

func Ascii(w http.ResponseWriter, r *http.Request) {
	// Retrieve form values from the HTTP request
	banner := r.PostFormValue("banners")
	input := r.PostFormValue("text")
	if r.Method != "POST" {
		functions.MessageError(w, r, http.StatusMethodNotAllowed, "Method Not Allowed") // Handle the error and return
		return
	}

	// Check if the 'banners' parameter is valid
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		functions.MessageError(w, r, http.StatusBadRequest, "Bad Request Invalid Banner") // Handle the error and return
		return
	}
	input2 := functions.CheckInput(input)

	// Call functions to get banner and process input
	getBanner := functions.GetBanner(banner)
	output := functions.ReadInput(input2, getBanner)

	// Parse the HTML template file
	tmpl, err := template.ParseFiles("templates/ascii-art.html")
	if err != nil {
		functions.MessageError(w, r, http.StatusInternalServerError, "Template parsing error")
		return
	}

	// Execute the template with the processed output
	err = tmpl.Execute(w, output)
	if err != nil {
		functions.MessageError(w, r, http.StatusInternalServerError, "Template execution error")
		return
	}
}
