package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}


func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "OK")
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/health", healthHandler)
    port := ":8080"
    log.Printf("Starting server on port %s\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}
