package main

import (
	"net/http"
	"testing"
)

func TestJulienschmidtParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("julienschmidt"))
}

func BenchmarkJulienschmidtSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("julienschmidt"), req)
}

func BenchmarkJulienschmidtParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher?name=gopher", nil)
	benchRequest(b, getHandler("julienschmidt"), req)
}
