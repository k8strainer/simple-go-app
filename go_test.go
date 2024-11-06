package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatalf("Failed to create request: %v", err)
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
