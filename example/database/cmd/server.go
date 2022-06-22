package main

import (
	"database/handlers"
	"fmt"
	"net/http"
)

func main() {
	port := ":4000"

	http.HandleFunc("/users", handlers.Users)
	http.HandleFunc("/books", handlers.Books)

	fmt.Println("server ishlayabdi", port)
	http.ListenAndServe(port, nil)
}

// func Books(rw http.ResponseWriter, r *http.Request) {
// 	rw.Header().Set("Content-type", "application/json")

// 	e := json.NewEncoder(rw)

// }
