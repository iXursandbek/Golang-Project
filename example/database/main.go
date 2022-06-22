package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Order struct {
	Id       int
	Name     string
	Statuses []int64
}

type OrderResponse struct {
	Id       int
	Name     string
	Statuses []string
}

func main() {
	conStr := "host=localhost port=5432 dbname=delivery user=postgres password=x3127106 sslmode=disable"
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		fmt.Println("bazaga ulanishda xatolik")
		panic(err)
	}
	defer db.Close()

	order := Order{Name: "KFC", Statuses: []int64{1, 3}}

	var id int64
	err = db.QueryRow(`INSERT INTO orders (name) VALUES $1 returning id`, order.Name).Scan(&id)
	if err != nil {
		panic(err)
	}

	for _, v := range order.Statuses {
		_, err = db.Exec(`INSERT INTO order_statuses (order_id, status_id) VALUES ($1,$2)`, id, v)
		if err != nil {
			panic(err)
		}
	}

	// rows := db.QueryRow(`SELECT o.id, o.name, s.name
	// 					FROM order_statuses os
	// 					JOIN orders o ON o.id = os.order_id
	// 					JOIN statuses s ON s.id = os.status_id`)

	newOrder := OrderResponse{}
	row := db.QueryRow(`SELECT id, name FROM orders WHERE id = $1`, id)
	err = row.Scan(
		&newOrder.Id,
		&newOrder.Name,
	)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT s.name FROM order_statuses os
							 JOIN statuses s os.status_id = s.id 
							 WHERE os.order_id = $1`, newOrder.Id)

	for rows.Next() {
		var statusName string
		err = rows.Scan(
			&statusName,
		)
		if err != nil {
			panic(err)
		}

		newOrder.Statuses = append(newOrder.Statuses, statusName)

	}

	fmt.Println(newOrder)

}
