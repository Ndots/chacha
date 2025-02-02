package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	//"structs"

	"chacha/backend/database"
	"chacha/backend/models"
	//"chacha/backend/utils"
)

type Director struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Position string `json:"position"`
}

func RegisterBusiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get user ID from context (assuming you've set it in auth middleware)
	userID := r.Context().Value("user_id").(int)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse request body
	var req struct {
		Name      string     `json:"name"`
		Address   string     `json:"address"`
		Type      string     `json:"type"`
		Email     string     `json:"email"`
		Phone     string     `json:"phone"`
		Directors []Director `json:"directors"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db := database.ConnectDB()
	defer db.Close()

	newBusiness := &models.Business{
		Name:      req.Name,
		Address:   req.Address,
		Type:      req.Type,
		Email:     req.Email,
		Phone:     req.Phone,
		Status:    "pending",
		UserID:    userID,
		Directors: []models.Director{
			{
				Name:     req.Directors[0].Name,
				Email:    req.Directors[0].Email,
				Position: req.Directors[0].Position,
			},
		},
	}

	if err := models.CreateBusiness(db, newBusiness); err != nil {
		log.Printf("Error creating business: %v", err)
		http.Error(w, "Failed to create business", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Business registered successfully",
		"business": newBusiness,
	})
}
