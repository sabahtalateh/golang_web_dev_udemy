package main

import (
	"database/sql"
	"io"
	"io/ioutil"
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
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.HandleFunc("/amigos", amigos)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from AWS")
}

func ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	io.WriteString(w, string(body))
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
