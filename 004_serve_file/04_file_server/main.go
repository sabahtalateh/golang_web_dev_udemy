package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/doggo", doggoPage)
	http.ListenAndServe(":8081", nil)
}

func doggoPage(w http.ResponseWriter, r *http.Request) {
	body := `<html><h1>DOGG</h1><img src="/doggo.jpg"/></html>`
	io.WriteString(w, body)
}
