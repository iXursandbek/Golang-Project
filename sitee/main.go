package main

import (
	"net/http"

	"example.com/sitee/controllers/homecontroller"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homecontroller.Index)
	http.HandleFunc("/home", homecontroller.Index)
	http.HandleFunc("/home/index", homecontroller.Index)

	http.ListenAndServe(":3000", nil)
}
