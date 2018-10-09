package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	// http.ListenAndServe
	err := http.ListenAndServeTLS(":10443", "localhost.crt", "localhost.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from https yoba")
}
