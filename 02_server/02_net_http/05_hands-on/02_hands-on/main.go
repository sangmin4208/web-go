package main

import (
	"fmt"
	"html/template"
	"net/http"
)


var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	http.HandleFunc("/dog/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Happy Dog")
	})
	http.HandleFunc("/me/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.gohtml", "Nacho")
	})
	http.ListenAndServe(":8080", nil)
}