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
	err := tpl.ExecuteTemplate(os.Stdout,"wisdom.tmpl",  `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}