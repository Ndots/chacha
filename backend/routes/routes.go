package routes

import (
	"chacha/backend/handlers"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	// TODO: Add actual registration logic here
	// For now, return a mock response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: Add actual authentication logic here
	// For now, return a mock response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": "mock-token",
	})
}

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// Auth endpoints
	router.HandleFunc("/api/register", handleRegister).Methods("POST")
	router.HandleFunc("/api/login", handleLogin)

	// Business endpoints
	router.HandleFunc("/api/business", handlers.RegisterBusiness).Methods("POST")

	// Application endpoints
	router.HandleFunc("/api/application", handlers.UpdateApplication).Methods("PUT")

	// Dashboard endpoint (for admin & user)
	router.HandleFunc("/api/dashboard", handlers.GetDashboardStats).Methods("GET")

	// Partner endpoint
	router.HandleFunc("/api/partner/application", handlers.PartnerApproveReject).Methods("PUT")

	return router
}
