package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `The Neeraj srivastava is assigning value to the vairiable.`)
	if err != nil {
		log.Fatalln(err)
	}
}
