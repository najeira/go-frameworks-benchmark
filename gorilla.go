package main

import (
	"fmt"
	"net/http"

	gorilla "github.com/gorilla/mux"
)

func init() {
	calcMem("gorilla", initGorilla)
}

func initGorilla() {
	h := gorilla.NewRouter()
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, World")
	})
	h.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s", gorilla.Vars(r)["name"])
	})
	registerHandler("gorilla", h)
}
