package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:rootroot@tcp(db.cofmeefxhubo.us-east-2.rds.amazonaws.com:3306)/db?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.ListenAndServe(":8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	_, e := io.WriteString(w, "at index")
	check(e)
}

func amigos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT name FROM amigos;`)
	if err != nil {
		log.Fatal(err)
	}

	var s, name string
	s = "RECEIVED RECORDS:\n"

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		s += name + "\n"
	}
	io.WriteString(w, s)
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customers (name VARCHAR(20));`)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, "CREATED TABLE customers", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customers VALUES('James')`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "INSERTED RECORDS", n)
}
