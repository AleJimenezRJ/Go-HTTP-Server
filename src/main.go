package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler) // Register the handler for the root path
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil) // Start the server
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
