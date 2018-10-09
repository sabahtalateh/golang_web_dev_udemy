package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Book stuct
type Book struct {
	isbn   string
	title  string
	author string
	price  float64
}

func main() {
	db, err := sql.Open("postgres", "postgres://usr:pwd@localhost:15432/mydb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	isbn := "999-1716216299"
	title := "Ivan Gog and Petr Gogog"
	author := "Petr Gogoge"
	price := 78.00

	res, err := db.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)", isbn, title, author, price)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
