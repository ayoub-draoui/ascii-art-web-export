package main

import (
	"fmt"
	"net/http"

	"functions/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/ascii-art", controllers.Ascii)
	http.HandleFunc("/export", controllers.Export)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("http://localhost:8084")
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		fmt.Println("Error Connected")
		return
	}
}
