package main

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatal(err)
		}
		c = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: false,
		}
		http.SetCookie(w, c)
	}
}
