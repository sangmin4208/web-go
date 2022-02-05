package main

import (
	"io"
	"net/http"
)

type nacho int

func (n nacho) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "dog")
	case "/cat":
		io.WriteString(w, "cat")
	}
}

func main(){
	var n nacho
	http.ListenAndServe(":8080", n)
}