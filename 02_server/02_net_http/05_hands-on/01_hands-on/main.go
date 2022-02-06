package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	http.HandleFunc("/dog/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Happy Dog")
	})
	http.HandleFunc("/me/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Nacho")
	})
	http.ListenAndServe(":8080", nil)
}