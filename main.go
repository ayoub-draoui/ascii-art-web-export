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

func index(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "index.html")
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, output)
}

func ascii(w http.ResponseWriter, r *http.Request) {
	// f := r.PostFormValue("banners")
	// v := r.PostFormValue("text")
	// w.Write([]byte("<h1>hello</h1>" + f + " " + v))

	// tmpl, _ := template.ParseFiles("index.html")
	banner := r.PostFormValue("banners")
	// fmt.Println(banner)
	input := r.PostFormValue("text")
	// banner = r.FormValue("banners")
	getBanner := functions.GetBanner(banner)
	output = functions.ReadInput(input, getBanner)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	// fmt.Println(output)

	// w.Write([]byte(output))

	// tmpl.Execute(w, output)
}

func main() {
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./style"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/ascii", ascii)
	// http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./images"))))

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
