package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(index)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", status, http.StatusOK)
	}

	expected := "Hello, World"
	if responseRecorder.Body.String() != expected {
		t.Errorf("Unexpected body: got %v, want %v", responseRecorder.Body.String(), expected)
	}
}
