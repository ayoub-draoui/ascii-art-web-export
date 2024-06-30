package functions

import (
	"net/http"
	"text/template"
)

var (
	tmpl   *template.Template
	output string
)

type Data struct {
	ErrNum int
	ErrTxt string
}

func HandlerError(w http.ResponseWriter, r *http.Request, d *Data) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	w.WriteHeader(d.ErrNum)
	tmpl.ExecuteTemplate(w, "error.html", d)
}
