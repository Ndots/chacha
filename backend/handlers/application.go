package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"chacha/backend/database"
	"chacha/backend/models"
	"chacha/backend/utils"
)

func UpdateApplication(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	applicationIDStr := r.URL.Query().Get("applicationId")
	if applicationIDStr == "" {
		http.Error(w, "Missing applicationId", http.StatusBadRequest)
		return
	}
	applicationID, err := strconv.Atoi(applicationIDStr)
	if err != nil {
		http.Error(w, "Invalid applicationId", http.StatusBadRequest)
		return
	}

	var req struct {
		Status string `json:"status"`
		Reason string `json:"reason"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := models.UpdateApplicationStatus(db, applicationID, req.Status, req.Reason); err != nil {
		http.Error(w, "Error updating application", http.StatusInternalServerError)
		return
	}

	utils.Log("Application " + strconv.Itoa(applicationID) + " updated to " + req.Status)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Application updated"})
}
