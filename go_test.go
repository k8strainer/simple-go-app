package main

import (
    "io"
    "net/http"
    "testing"
    "time"
)

func TestMainFunction(t *testing.T) {
    // Start the server in a goroutine
    go func() {
        main()
    }()

    // Give the server a moment to start
    time.Sleep(100 * time.Millisecond)

    // Perform a request to the root endpoint
    resp, err := http.Get("http://localhost:8080")
    if err != nil {
        t.Fatalf("Failed to reach server: %v", err)
    }
    defer resp.Body.Close()

    // Check if the response is as expected
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status OK, got %v", resp.StatusCode)
    }

    // Check the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }

    expected := "Hello, World!\n"
    if string(body) != expected {
        t.Errorf("Expected response body %q, got %q", expected, string(body))
    }
}
