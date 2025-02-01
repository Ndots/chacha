package handlers

import (
	"encoding/json"
	"net/http"

	"chacha/backend/database"
	"chacha/backend/models"
	"chacha/backend/utils"
)

func RegisterBusiness(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	defer db.Close()

	var business models.Business
	if err := json.NewDecoder(r.Body).Decode(&business); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := models.CreateBusiness(db, &business); err != nil {
		http.Error(w, "Error registering business", http.StatusInternalServerError)
		return
	}

	utils.Log("Business registered: " + business.Name)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(business)
}
