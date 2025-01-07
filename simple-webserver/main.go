package main

import (
	"fmt"
	"log"
	"net/http"
)

// function to handle form requests
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// function to handle the hello requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported ", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

// main entry point for a web server
func main() {
	// init the fileServer for static files
	fileServer := http.FileServer(http.Dir("./static"))

	// on root path load the index file
	http.Handle("/", fileServer)

	// endpoint to handle form and hello
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// start the server which listen the requests mn port 8080
	fmt.Printf("Starting server at port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
