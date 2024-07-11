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
	fmt.Println("http://localhost:8083")
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		fmt.Println("Error Connected")
		return
	}
}
