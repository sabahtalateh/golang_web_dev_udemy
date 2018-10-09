package main

import (
	"database/sql"
	"fmt"
	"log"

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
	fmt.Println("You connected")

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		b := Book{}
		err := rows.Scan(&b.isbn, &b.title, &b.author, &b.price)
		if err != nil {
			log.Println(err)
		}
		books = append(books, b)
	}
	fmt.Printf("%+v", books)
}
