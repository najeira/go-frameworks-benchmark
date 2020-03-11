package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	calcMem("julienschmidt", initJulienschmidt)
}

func initJulienschmidt() {
	r := httprouter.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, World")
	})
	r.GET("/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s", ps.ByName("name"))
	})
	registerHandler("julienschmidt", r)
}
