package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
	appengine.Main()
}

func index(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
