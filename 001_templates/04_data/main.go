package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	fname string
	lname string
}

func main() {
	tpl, err := template.ParseFiles("templates/1.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	p := person{
		fname: "Ivan",
		lname: "Gog",
	}
	tpl.Execute(os.Stdout, p)
	tpl.Execute(os.Stdout, 42)
}
