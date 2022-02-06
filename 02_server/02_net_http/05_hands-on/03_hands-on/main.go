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
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(myName))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, "{name: Happy Dog}")
}
func myName(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "Nacho")
}