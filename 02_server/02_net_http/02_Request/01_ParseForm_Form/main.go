package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)


type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(r.Form)
	fmt.Println(r.PostForm)
	fmt.Println(r.RequestURI)
	tpl.ExecuteTemplate(w, "tpl.gohtml", r.Form)
}

var tpl *template.Template

func init (){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main(){
	var d hotdog
	err := http.ListenAndServe(":8080", d)
	if err != nil{
		panic(err)
	}
}