package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// adds entire templates folder to be parsed 
	tpl = template.Must(template.ParseGlob("template/*")) 

	//Singular file
	// tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main()  {
	//data
	err := tpl.ExecuteTemplate(os.Stdout, "data.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
	//variable
	err = tpl.ExecuteTemplate(os.Stdout, "variable.gohtml", "Focus on being your best")
	if err != nil {
		log.Fatalln(err)
	}
}
