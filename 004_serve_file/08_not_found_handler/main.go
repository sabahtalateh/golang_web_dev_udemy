package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", f)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", nil)
}

func f(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprint(w, "Go look to your terminal")
}
