package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	instance := &handler{}
	http.Handle("/hello", instance)
	http.HandleFunc("/hello_func", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from func, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
