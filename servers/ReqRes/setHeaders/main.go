package main

import (
	"fmt"
	"net/http"
)

type thing int

func (t thing) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("My key", "this is my key")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Println(w, "<h1>Any code you want to add</h1>")
}

func main() {
	var th thing
	http.ListenAndServe(":8080", th)
}
