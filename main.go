package main

import (
	"fmt"
	"html/template"
	"net/http"

	"functions/functions"
)

var (
	tmpl   *template.Template
	output string
)

type Data struct {
	ErrNum int
	ErrTxt string
}

func index(w http.ResponseWriter, r *http.Request) {
	d := Data{}
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	// ERROR 404 Page Not Found
	if r.URL.Path != "/" {
		d.ErrNum = http.StatusNotFound
		d.ErrTxt = "Page Not Found"
		HandlerError(w, r, &d)
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
		HandlerError(w, r, &d)
		return
	}
}

func ascii(w http.ResponseWriter, r *http.Request) {
	// Initialize a struct to handle potential errors
	d := Data{}

	// Retrieve form values from the HTTP request
	banner := r.PostFormValue("banners")
	input := r.PostFormValue("text")

	// Check if the 'banners' parameter is valid
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		d.ErrNum = http.StatusInternalServerError
		d.ErrTxt = "Invalid banner type"
		HandlerError(w, r, &d) // Handle the error and return
		return
	}
	if !functions.CheckInput(input) {
		d.ErrNum = http.StatusBadRequest
		d.ErrTxt = "Bad Request"
		HandlerError(w, r, &d)
		return
	}

	// Call functions to get banner and process input
	getBanner := functions.GetBanner(banner)
	output := functions.ReadInput(input, getBanner)

	// Parse the HTML template file
	tmpl, err := template.ParseFiles("templates/ascii-art.html")
	if err != nil {
		d.ErrNum = http.StatusInternalServerError
		d.ErrTxt = "Template parsing error"
		HandlerError(w, r, &d) // Handle the error and return
		return
	}

	// Execute the template with the processed output
	err = tmpl.Execute(w, output)
	if err != nil {
		d.ErrNum = http.StatusInternalServerError
		d.ErrTxt = "Template execution error"
		HandlerError(w, r, &d) // Handle the error and return
		return
	}
}

func HandlerError(w http.ResponseWriter, r *http.Request, d *Data) {
	w.WriteHeader(d.ErrNum)
	tmpl.ExecuteTemplate(w, "error.html", d)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ascii-art", ascii)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./style"))))

	fmt.Println("http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
