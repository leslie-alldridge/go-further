package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main(){
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"Math111", "Intro to Math", "3"},
				course{"Sci101", "Intro to Sci", "3"},
				course{"Comp101", "Intro to Comp", "3"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"Math112", "Intermediate Math", "3"},
				course{"Sci202", "Intermediate Sci", "3"},
				course{"Comp202", "Intermediate Comp", "3"},
			},
		},
	}
	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}