package main

import (
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

func TestHelloHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(helloHandler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := "Hello, World!\n"
    if rr.Body.String() != expected {
        t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

func TestHealthHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(healthHandler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := "OK\n"
    if rr.Body.String() != expected {
        t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

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

    // Test /health endpoint as part of main function test
    resp, err = http.Get("http://localhost:8080/health")
    if err != nil {
        t.Fatalf("Failed to reach health endpoint: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status OK, got %v", resp.StatusCode)
    }

    body, err = io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("Failed to read response body: %v", err)
    }

    expected = "OK\n"
    if string(body) != expected {
        t.Errorf("Expected response body %q, got %q", expected, string(body))
    }
}
