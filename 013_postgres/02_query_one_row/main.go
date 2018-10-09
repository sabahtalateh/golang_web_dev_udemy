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

	isbn := "172-1716216273"
	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)
	if err != nil {
		log.Println(err)
	}
	var b Book

	err = row.Scan(&b.isbn, &b.author, &b.title, &b.price)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("Not Found")
		return
	case err != nil:
		panic(err)
	}

	fmt.Printf("%+v", b)
}
