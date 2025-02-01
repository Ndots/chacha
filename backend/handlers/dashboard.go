package handlers

import (
	"encoding/json"
	"net/http"

	"chacha/backend/database"
	"chacha/backend/models"
)

func GetDashboardStats(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	stats, err := models.GetDashboardStats(db)
	if err != nil {
		http.Error(w, "Error retrieving dashboard stats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
