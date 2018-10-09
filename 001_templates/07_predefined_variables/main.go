package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

type d struct {
	Xs    []string
	Users []user
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	u1 := user{
		Name:  "Buddha",
		Motto: "Go meditate",
		Admin: false,
	}

	u2 := user{
		Name:  "Gandhi",
		Motto: "Go eat some Gandghi cakes",
		Admin: true,
	}

	xs := []string{"zero", "one", "two", "three", "four"}
	users := []user{u1, u2}

	err := tpl.Execute(os.Stdout, d{Xs: xs, Users: users})
	if err != nil {
		log.Fatal(err)
	}
}
