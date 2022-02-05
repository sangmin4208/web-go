package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"cat cat cat")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/dog/",hotdog(0))
	mux.Handle("/cat",hotcat(0))

	http.ListenAndServe(":8080",mux)
}