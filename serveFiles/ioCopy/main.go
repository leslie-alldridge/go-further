package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/bella.jpg", catPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/bella.jpg">
	`)
}

func catPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("bella.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
