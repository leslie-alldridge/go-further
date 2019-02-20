package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", bar)
	http.HandleFunc("/me", me)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo")
}

func bar(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("./templates/dog.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "dog.gohtml", "McLeod")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Leslie")
}
