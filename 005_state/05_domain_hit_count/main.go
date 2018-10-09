package main

import (
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", page)
	http.ListenAndServe(":8081", nil)
}

func page(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("count")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "count",
			Value: "0",
		}
	}
	count, err := strconv.Atoi(c.Value)
	if err != nil {
		count = 0
	}
	count++
	c.Value = strconv.Itoa(count)
	http.SetCookie(w, c)
}
