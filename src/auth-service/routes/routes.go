package routes

import (
	"auth-service/controllers"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	r.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")
	return r
}