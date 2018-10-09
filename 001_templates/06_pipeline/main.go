package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

type person struct {
	fname string
	lname string
}

var tpl *template.Template

var fm = template.FuncMap{
	"fdateMDY": mdy,
}

func mdy(t time.Time) string {
	return t.Format("01-02-2006")
	// return t.Format(time.ANSIC)
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
