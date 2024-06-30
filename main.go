package main

import (
	"fmt"
	"net/http"

	"functions/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/ascii-art", controllers.Ascii)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./style"))))
	// http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	fmt.Println("http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
