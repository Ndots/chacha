package handlers

import (
	"chacha/backend/models"
	"encoding/json"
	"net/http"
	"os"
	//"log"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// If JWT_SECRET is not set in your environment, you can temporarily set a default for development:
//var jwtSecret = []byte("my-secret123")

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// RegisterUser creates a new user and persists it to the database.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: Add your actual registration logic here

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

// LoginUser authenticates the user and returns a JWT token upon successful login.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Set headers
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: Add your actual authentication logic here

	// For now, return a success response
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"token":   "your-jwt-token-here",
	})
}
