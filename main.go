package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// Log the request to stdout
	fmt.Println("Received request for /:", r.Method, r.URL.Path)

	// Send a response to the client
	fmt.Fprintln(w, "Hello from the main path!")
}

func otherHandler(w http.ResponseWriter, r *http.Request) {
	// Log the request to stdout
	fmt.Println("Received request for /other:", r.Method, r.URL.Path)

	// Send a response to the client
	fmt.Fprintln(w, "Hello from the other path!")
}

func main() {
	// Define the server's address
	addr := ":8080"

	// Register the handler functions for the corresponding URL patterns
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/other", otherHandler)

	// Start the server and log any errors
	fmt.Printf("Starting server on %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
