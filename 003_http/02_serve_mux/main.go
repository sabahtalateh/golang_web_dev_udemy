package main

import (
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DD OOO GG!! IT IS DOOOOOOG"))
}

type hotcat int

func (hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("cat"))
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)
	http.ListenAndServe(":8081", mux)
}
