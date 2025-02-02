package handlers

import (
	"chacha/backend/database"
	"chacha/backend/models"
	"chacha/backend/utils"
	"encoding/json"
	"net/http"
	//"os"
	"time"

	"github.com/golang-jwt/jwt"
	//"log"
)

//var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
// If JWT_SECRET is not set in your environment, you can temporarily set a default for development:
var jwtSecret = []byte("my-secret123")

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
	db := database.ConnectDB()
	defer db.Close()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error processing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	if err := models.CreateUser(db, &user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	utils.Log("User registered: " + user.Email)

	// Remove the password before returning the response
	user.Password = ""
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// LoginUser authenticates the user and returns a JWT token upon successful login.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByEmail(db, loginReq.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(loginReq.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Create a JWT token valid for 72 hours.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Remove password before sending user data in response.
	user.Password = ""
	res := LoginResponse{
		Token: tokenString,
		User:  *user,
	}
	json.NewEncoder(w).Encode(res)
}
