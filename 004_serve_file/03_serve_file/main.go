package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/doggo", doggoPage)
	http.HandleFunc("/doggo.jpg", doggoImage)

	http.ListenAndServe(":8081", nil)
}

func doggoPage(w http.ResponseWriter, r *http.Request) {
	body := `<html><h1>DOGG</h1><img src="/doggo.jpg"/></html>`
	io.WriteString(w, body)
}

func doggoImage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "doggo.jpg")
}
