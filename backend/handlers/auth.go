package handlers

import (
	"encoding/json"
	"net/http"

	"chacha/backend/database"
	"chacha/backend/models"
	"chacha/backend/utils"
	//"log"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

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

	// For development, we return a dummy token.
	token := "dummy-token"
	res := LoginResponse{
		Token: token,
		User:  *user,
	}
	json.NewEncoder(w).Encode(res)
}
