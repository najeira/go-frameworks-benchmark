package main

import (
	"net/http"
	"testing"
)

func BenchmarkHttp(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("http"), req)
}
