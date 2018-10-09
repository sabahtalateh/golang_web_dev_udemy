package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8081", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:  "my-cookie",
		Value: "my-value",
	}
	http.SetCookie(w, &c)
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		io.WriteString(w, "Cookie not found")
		return
	}
	io.WriteString(w, fmt.Sprintf("%+v", c))
}

func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != http.ErrNoCookie {
		c.MaxAge = -1
		http.SetCookie(w, c)
	}
}
