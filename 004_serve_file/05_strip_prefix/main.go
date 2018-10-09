package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/dog", doggoPage)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8081", nil)
}

func doggoPage(w http.ResponseWriter, r *http.Request) {
	body := `<html><h1>DOGG</h1><img src="/resources/doggo.jpg"/></html>`
	io.WriteString(w, body)
}
