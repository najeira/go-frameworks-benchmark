package main

import (
	"net/http"
	"testing"
)

func BenchmarkGocraft(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("gocraft"), req)
}
