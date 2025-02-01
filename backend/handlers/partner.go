package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"chacha/backend/database"
	"chacha/backend/models"
	"chacha/backend/utils"
)

func PartnerApproveReject(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	var req struct {
		ApplicationID int    `json:"application_id"`
		Status        string `json:"status"`
		Reason        string `json:"reason"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := models.UpdateApplicationStatus(db, req.ApplicationID, req.Status, req.Reason); err != nil {
		http.Error(w, "Error updating application", http.StatusInternalServerError)
		return
	}

	utils.Log("Partner updated application " + strconv.Itoa(req.ApplicationID))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Application status updated by partner"})
}
