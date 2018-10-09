package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/doggo.jpg", dogggo)
	http.HandleFunc("/doggo", doggoPage)
	http.ListenAndServe(":8081", nil)
}

func doggoPage(w http.ResponseWriter, r *http.Request) {
	resp := `<h1>Hello</h1><img src="/doggo.jpg">`
	io.WriteString(w, resp)
}

func dogggo(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("doggo.jpg")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
