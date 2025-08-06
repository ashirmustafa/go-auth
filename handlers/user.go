package handlers

import (
	"encoding/json"
	"fmt"
	"go-auth/config"
	"go-auth/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var input UserInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashed, err := HashPassword("mypassword123")
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Email:    input.Email,
		Password: hashed,
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created: %s", user.Email)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	result := config.DB.Find(&users)
	if result.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
