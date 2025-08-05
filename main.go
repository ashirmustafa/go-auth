package main

import (
    "go-auth/config"
    "go-auth/routes"
    "log"
    "net/http"

    "github.com/gorilla/handlers"
)

func main() {
    err := config.ConnectDB()
    if err != nil {
        log.Fatal(err)
    }

    routes.RegisterUserRoutes()

    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}), // allow all for dev
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )(http.DefaultServeMux)

    log.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", corsHandler)
}
