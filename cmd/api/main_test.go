package main

import (
	"io"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/hello?name=Nathan", nil)
	w := httptest.NewRecorder()

	HelloHandler(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if string(body) != "Hello, Nathan" {
		t.Fatalf(`expected "Hello, Nathan", got "%s"`, string(body))
	}
}
