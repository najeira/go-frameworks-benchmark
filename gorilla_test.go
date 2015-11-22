package main

import (
	"net/http"
	"testing"
)

func TestGorillaParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("gorilla"))
}

func BenchmarkGorillaSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("gorilla"), req)
}

func BenchmarkGorillaParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("gorilla"), req)
}
