package main

import (
	"net/http"
	"testing"
)

func BenchmarkGoji(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("goji"), req)
}
