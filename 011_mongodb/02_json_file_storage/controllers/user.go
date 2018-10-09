package controllers

import (
	"encoding/json"
	"fmt"
	"go_web_dev/011_mongodb/02_json_file_storage/models"
	"log"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/julienschmidt/httprouter"
)

// UserController is an empty struct to keep methods
type UserController struct {
	Session map[string]models.User
}

// NewUserController creates User Controller
func NewUserController(s map[string]models.User) *UserController {
	return &UserController{
		Session: s,
	}
}

// CreateUser is for user creation
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	uid, _ := uuid.NewV4()
	u.ID = uid.String()

	uc.Session[u.ID] = u
	models.StoreUsers(uc.Session)

	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// GetUser returns a user for a given id
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	m := models.LoadUsers()
	u, ok := m[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		log.Println(err)
	}
}

// DeleteUser deletes given user
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	m := models.LoadUsers()

	delete(m, id)

	models.StoreUsers(m)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted")
}
