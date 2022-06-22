package handlers

import (
	"database/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var users = []models.User{
	{Name: "Jahongir", Age: 12},
	{Name: "Abdulaziz", Age: 23},
	{Name: "Muhammadqodir", Age: 18},
	{Name: "Kamron", Age: 20},
	{Name: "Husniddin", Age: 21},
}

var books = []models.Book{
	{Name: "Baxtiyor oila", Price: 1200},
	{Name: "Goliblik qonuniytlari", Price: 2300},
	{Name: "Omomning maniken qizi", Price: 1200},
	{Name: "Qo'rqma", Price: 3400},
	{Name: "molxona", Price: 2333},
}

func Users(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var user models.User
	json.Unmarshal(body, &user)

	fmt.Println(user)
	users = append(users, user)

	e := json.NewEncoder(rw)
	e.Encode(users)
}

func Books(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")

	e := json.NewEncoder(rw)
	e.Encode(books)
}
