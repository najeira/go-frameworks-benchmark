package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

func init() {
	calcMem("martini", initMartini)
}

func initMartini() {
	r := martini.NewRouter()
	m := martini.New()
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	h := &martini.ClassicMartini{m, r}
	h.Get("/", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		return "Hello, World"
	})
	h.Get("/:name", func(w http.ResponseWriter, params martini.Params) string {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		return fmt.Sprintf("Hello, %s", params["name"])
	})
	registerHandler("martini", h)
}
