package main

import (
	"fmt"
	"net/http"

	"github.com/xujiajun/gorouter"
)

func init() {
	calcMem("gorouter", initGorouter)
}

func initGorouter() {
	r := gorouter.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, World")
	})
	r.GET("/:name", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s", gorouter.GetParam(r, "name"))
	})
	registerHandler("gorouter", r)
}
