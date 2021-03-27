package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	testServer := httptest.NewServer(CreateServer())
	defer testServer.Close()

	response, err := http.Get(fmt.Sprintf("%s/ping", testServer.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", response.StatusCode)
	}

	contentTypeHeader, ok := response.Header["Content-Type"]
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}
	if contentTypeHeader[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", contentTypeHeader[0])
	}
}
