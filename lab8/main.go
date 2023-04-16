package main

import (
	"fmt"
	"net/http"
	"time"
)

//requests to the root URL path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my website!")
}
//requests to the /about URL path
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a brief description of my website.")
}
//my func:
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

// registering the home and about handlers with the default HTTP server:
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
// starting the server on port 8080
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
	// creating a new server with custom configuration
    srv := &http.Server{
        Addr:         ":8080",
        Handler:      nil, // using the default ServeMux
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    // registering a handler function for the root URL
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    // start the server and listen for incoming requests
    if err := srv.ListenAndServe(); err != nil {
        panic(err)
    }
}
