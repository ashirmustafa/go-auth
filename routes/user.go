package routes

import (
	"net/http"
	"go-auth/handlers"
)

func RegisterUserRoutes() {
	http.HandleFunc("/create-account", handlers.CreateUser) 
}
