package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    port := ":8080"
    log.Printf("Starting server on port %s\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}
