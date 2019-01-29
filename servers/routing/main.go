package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog doggy")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat catty")
}

func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)

	http.ListenAndServe(":8080", nil)
}
