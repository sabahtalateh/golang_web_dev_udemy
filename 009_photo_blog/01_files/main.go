package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/demo", demo)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":8081", nil)
}

func demo(w http.ResponseWriter, r *http.Request) {
	url := "https://srg.zoz:4443/oauth/v2/auth?client_id=2_4h0uilbo77s4gkwccws08s0wg44ww0wcgko4cksw48sgokcgss&redirect_uri=http%3A//localhost%3A8081/demo&response_type=code"
	link := fmt.Sprintf(`<h1>Auth with auth code</h1><a href="%s">%s</a>`, url, url)
	io.WriteString(w, link)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("nf")
		if err != nil {
			log.Println(err)
		}
		defer mf.Close()
		// sha from filename
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		// create new file
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fname)
		nf, err := os.Create(path)
		if err != nil {
			log.Println(err)
		}
		defer nf.Close()
		// copy file
		mf.Seek(0, 0)
		io.Copy(nf, mf)
		// add filename to cookie
		appendValue(w, c, fname)
	}
	strings := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", strings[1:])
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
