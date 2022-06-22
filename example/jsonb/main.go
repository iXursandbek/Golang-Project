package main

import (
	"database/sql"

	// "encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastNAme  string `json:"last_n_ame"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	// Phones    []Phone `json:"phones"`
}

type Phone struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "x3127106"
	dbname   = "fruitdb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	connect, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer connect.Close()

	err = connect.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	// user := User{
	// 	FirstName: "Elon",
	// 	LastNAme:  "Mask",
	// 	UserName:  "elonmask2",
	// 	Email:     "elon@maski.com",
	// 	Phones: []Phone{
	// 		{
	// 			Type:   "mobule",
	// 			Number: "1234444",
	// 		},
	// 		{
	// 			Type:   "Home",
	// 			Number: "3456",
	// 		},
	// 	},
	// }

	// query := `insert into users (first_name, last_name, username, email, phones) values ($1,$2,$3,$4,$5)`

	// phones, err := json.Marshal(user.Phones)
	// if err != nil {
	// 	panic(err)
	// }

	// row, err := connect.Exec(query, user.FirstName, user.LastNAme, user.UserName, user.Email, phones)
	// if err != nil {
	// 	fmt.Println("databasaga ulanishada xato")
	// }
	// fmt.Println(row.LastInsertId())

	user := User{}
	// var phoneString string
	row2 := connect.QueryRow(`select id, first_name, last_name, username, email from users where id = $1`, 1)
	err = row2.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastNAme,
		&user.UserName,
		&user.Email,
		// &user.Phones,
	)
	if err != nil {
		panic(err)
	}

	deleteQuery := connect.QueryRow(`delete from users where id = $1`, 3)
	err = deleteQuery.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastNAme,
		&user.UserName,
		&user.Email,
	)
	if err != nil {
		panic(err)
	}

	// err = json.Unmarshal([]byte(phoneString), &user.Phones)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(user)
}
