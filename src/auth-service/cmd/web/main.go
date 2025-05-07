package main 

import (
	"log"
	"net/http"
	"api-bengkel/auth-service/routes"
)

func main(){
	r := routes.InitRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))
}