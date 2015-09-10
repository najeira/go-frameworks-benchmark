package main

import (
	"github.com/go-martini/martini"
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
	h.Get("/", func() string {
		return responseText
	})
	registerHandler("martini", h)
}
