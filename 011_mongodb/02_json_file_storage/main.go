package main

import (
	"go_web_dev/011_mongodb/02_json_file_storage/controllers"
	"go_web_dev/011_mongodb/02_json_file_storage/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var users map[string]models.User

func main() {
	uc := controllers.NewUserController(getSession())

	r := httprouter.New()

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe(":8081", r)
}

func getSession() map[string]models.User {
	return models.LoadUsers()
}
