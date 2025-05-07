package services

import (
	"auth-service/utils"
	"auth-service/models"
	"errors"
	"net/http"
)

func Login(r *http.Request) (*models.User, error){
	email := r.FormValue("email")
	password := r.FormValue("password")

	if !utils.ValidateEmail(email){
		return nil, errors.New("Invalid email format")
	}

	user := &models.User{Email: email, Password: password}

	return user, nil
}