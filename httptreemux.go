package main

import (
	"fmt"
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
)

func init() {
	calcMem("httptreemux", initHttptreemux)
}

func initHttptreemux() {
	r := httptreemux.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, World")
	})
	r.GET("/:name", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s", params["name"])
	})
	registerHandler("httptreemux", r)
}
