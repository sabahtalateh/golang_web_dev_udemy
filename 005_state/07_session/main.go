package main

import (
	"html/template"
	"io"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string
	Fname    string
	Lname    string
	Password []byte
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionDuration int = 30

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8081", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	updateCookie(w, r)
	io.WriteString(w, "Hello")
}

func bar(w http.ResponseWriter, r *http.Request) {
	updateCookie(w, r)
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	s, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u, ok := dbUsers[s.un]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if u.Role != "007" {
		http.Error(w, "Only 007 can enter the bar", 401)
		return
	}

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	updateCookie(w, r)
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/bar", http.StatusSeeOther)
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")

		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", 400)
			return
		}
		pass := r.FormValue("password")
		gen, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		role := r.FormValue("role")
		u := user{
			Username: un,
			Fname:    fname,
			Lname:    lname,
			Password: gen,
			Role:     role,
		}
		dbUsers[un] = u

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	updateCookie(w, r)
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Denied", 401)
			return
		}
		pass := r.FormValue("password")

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(pass))
		if err != nil {
			http.Error(w, "Denied", 401)
			return
		}

		c, err := r.Cookie("session")
		sID, _ := uuid.NewV4()
		if err != nil {
			c = &http.Cookie{
				Name:  "session",
				Value: sID.String(),
			}
			http.SetCookie(w, c)
		}

		dbSessions[c.Value] = session{u.Username, time.Now()}
		http.Redirect(w, r, "/bar", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	updateCookie(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 2) {
		go cleanSessions()
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
