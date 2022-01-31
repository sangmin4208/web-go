package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){ 
	budda := sage {
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	err := tpl.Execute(os.Stdout, budda)
	if err != nil {
		log.Fatalln(err)
	}
}