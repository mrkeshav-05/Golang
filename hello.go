package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define a handler function for the root endpoint
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!") // Send "Hello, World!" as the response
    })

    // Start the server on port 8080
    fmt.Println("Server is running on http://localhost:8000")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
