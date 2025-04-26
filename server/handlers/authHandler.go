package handlers

import (
	"cosmic-gate-chat/models"
	"cosmic-gate-chat/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

// Authenticate an User Account
func AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest models.User
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userData, err := repositories.AuthUser(userRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error auth user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userData)
}
