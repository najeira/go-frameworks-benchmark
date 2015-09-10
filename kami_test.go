package main

import (
	"net/http"
	"testing"
)

func BenchmarkKami(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("kami"), req)
}
