package main

import (
	"fmt"
	"net/http"

	"github.com/celrenheit/lion"
)

func init() {
	calcMem("lion", initLion)
}

func initLion() {
	r := lion.New()
	r.GetFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, World")
	})
	r.GetFunc("/:name", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s", lion.Param(r, "name"))
	})
	registerHandler("lion", r)
}
