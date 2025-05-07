package controllers

import (
	"api-bengkel/auth-service/services"
	"auth-service/models"
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	user, err := services.Login(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}
}