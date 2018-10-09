package main

import (
	"io"
	"net/http"
	"os"
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
	f, err := os.Open("doggo.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}
