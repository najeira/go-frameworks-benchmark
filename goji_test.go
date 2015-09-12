package main

import (
	"net/http"
	"strings"
	"testing"
)

func TestGojiParam(t *testing.T) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	c, b, h := sendRequest(getHandler("goji"), req)
	if c != 0 && c != 200 {
		t.Errorf("invalid status code")
	}
	if !strings.Contains(string(b), "gopher") {
		t.Errorf("invalid body")
	}
	if h == nil || !strings.Contains(h["Content-Type"][0], "text/plain") {
		t.Errorf("invalid header")
	}
}

func BenchmarkGojiSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("goji"), req)
}

func BenchmarkGojiParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("goji"), req)
}
