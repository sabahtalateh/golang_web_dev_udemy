package main

import (
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DD OOO GG!! IT IS DOOOOOOG"))
}

func c(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("cat"))
}

func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)
	http.Handle("/hello", http.HandlerFunc(c))
	http.ListenAndServe(":8081", nil)
}
