package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	sages := []string{"Gandhi", "Buddha", "Jesus"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatal(err)
	}

}
