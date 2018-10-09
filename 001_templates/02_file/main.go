package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal("error while parsing template", err)
	}
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error while creating file", err)
	}
	t.Execute(nf, nil)
}
