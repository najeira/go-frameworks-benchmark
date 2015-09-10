package main

import (
	"net/http"
	"testing"
)

func BenchmarkMartini(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("martini"), req)
}
