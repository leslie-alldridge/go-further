package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type human struct {
	Location string
	Age int
}

type house struct {
	Street string
	Value int
}

type info struct {
	Person []human
	Home []house
}

func init() {
	// adds entire templates folder to be parsed 
	tpl = template.Must(template.ParseGlob("template/*")) 

	//Singular file being parsed
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
	//slice of string
	slice := []string{"Leslie", "Aaron", "Lisa"}
	err = tpl.ExecuteTemplate(os.Stdout, "sliceofstring.gohtml", slice)
	if err != nil {
		log.Fatalln(err)
	}
	//map
	maps := map[string]string{
		"Top": "North",
		"Bottom": "South",
		"Right": "East",
		"Left": "West",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "maps.gohtml", maps)
	if err != nil {
		log.Fatalln(err)
	}
	//struct
	leslie := human{
		Location: "Wellington",
		Age: 26,
	}
	err = tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", leslie)
	if err != nil {
		log.Fatalln(err)
	}
	//slice of structs
	mike := human{
		Location: "New Plymouth",
		Age: 30,
	}
	people := []human{leslie,mike}
	err = tpl.ExecuteTemplate(os.Stdout, "sliceofstruct.gohtml", people)
	if err != nil {
		log.Fatalln(err)
	}
	//complex
	h := house{
		Street: "Young Street",
		Value: 400,
	}
	h2 := house {
		Street: "Old Street",
		Value: 200,
	}
	houses := []house{h,h2}
	peeps := []human{leslie,mike}
	data := info{
		Person: peeps,
		Home: houses,
	}
	//can also do a composite literal (creating a struct on the fly)

	// data2 := struct {
	// 	Person []human
	// 	Home []house
	// }{
	// 	peeps,
	// 	houses,
	// }

	err = tpl.ExecuteTemplate(os.Stdout, "complex.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
