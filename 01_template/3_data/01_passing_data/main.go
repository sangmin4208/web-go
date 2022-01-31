package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init (){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
	err := tpl.ExecuteTemplate(os.Stdout,"hello.tmpl", "Vespa")
	if err != nil {
		log.Fatalln(err)
	}
}