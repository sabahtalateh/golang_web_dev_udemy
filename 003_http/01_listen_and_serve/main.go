package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(b))

	r.ParseForm()
	fname := r.Form.Get("fname")
	fmt.Println(fname)
	fmt.Println(r.Form)
	fmt.Println(r.Method)
	fmt.Println(r.Form)
	fmt.Println(r.URL)
	w.Header().Set("My-Header", "zopa")
	fmt.Fprintln(w, "HELLO")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8081", d)
}
