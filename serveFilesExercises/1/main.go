package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/bella.jpg", catPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("./templates/cat.gohtml")
	if err != nil {
		log.Fatalln("error with template", err)
	}

	err = tpl.ExecuteTemplate(w, "cat.gohtml", "Leslie")
	if err != nil {
		log.Fatalln("error executing the template", err)
	}
}

func catPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "bella.jpg")
}
