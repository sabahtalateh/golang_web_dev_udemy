package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8081", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at foo " + r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at bar " + r.Method)
	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusSeeOther)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
