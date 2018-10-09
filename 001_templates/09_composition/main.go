package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type course struct {
	Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

func (s semester) CountCourses() int {
	return len(s.Courses)
}

func (s semester) TakesArg(a int) int {
	return s.CountCourses() + a
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

func main() {

	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{Name: "Golang for puppies", Units: "4"},
				course{Name: "Java for kittens", Units: "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{Name: "C++ for dogs", Units: "99"},
			},
		},
	}

	fmt.Println(y.Fall.CountCourses())

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatal(err)
	}
}
