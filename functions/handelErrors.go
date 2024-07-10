package functions

import (
	"net/http"
	"text/template"
)

type ErrorResponse struct {
	ErrNum int
	ErrTxt string
}

func MessageError(w http.ResponseWriter, r *http.Request, code int, msg string) {
	msg_error := &ErrorResponse{ErrNum: code, ErrTxt: msg + "!"}
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	w.WriteHeader(code)
	tmpl.ExecuteTemplate(w, "error.html", msg_error)
}
