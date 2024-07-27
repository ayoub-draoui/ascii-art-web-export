package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"functions/functions"
)

var dataglobal string

func Ascii(w http.ResponseWriter, r *http.Request) {
	// Retrieve form values from the HTTP request
	banner := r.PostFormValue("banners")
	input := r.PostFormValue("text")
	if r.Method != "POST" {
		functions.MessageError(w, r, http.StatusMethodNotAllowed, "Method Not Allowed") // Handle the error and return
		return
	}
	count := 0
	for _, char := range input {
		if char == '\r' {
			continue
		}
		count++
	}
	if count > 500 || count == 0 {
		functions.MessageError(w, r, http.StatusMethodNotAllowed, "Text is too long")
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

	// data reload
	dataglobal = output

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

func ExportASCIIArt(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && dataglobal != "" {

		// Set the response headers
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")

		// Set the Content-Length header
		w.Header().Set("Content-Length", strconv.Itoa(len(dataglobal)))

		// Write the ASCII art to the response
		_, err := w.Write([]byte(dataglobal))
		if err != nil {
			functions.MessageError(w, r, http.StatusInternalServerError, "Template execution error")
			return
		}
	} else {
		functions.MessageError(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
}
