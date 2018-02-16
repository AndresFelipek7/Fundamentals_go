package main

import (
	"net/http"
)

func main() {
	//Make Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	//Start the server
	http.ListenAndServe(":3000",nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact Page!!!!"))
}