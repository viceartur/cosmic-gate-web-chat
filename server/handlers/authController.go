package handlers

import (
	"cosmic-gate-chat/models"
	"cosmic-gate-chat/services"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Password Encryption
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Check whether the password provided is correct
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Authenticate an User Account
func AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest models.User
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundUser, err := services.GetUserByEmail(userRequest.Email)
	if foundUser.Email == "" || err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	ok := CheckPasswordHash(userRequest.Password, foundUser.Password)
	if !ok {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	userData := &models.User{
		ID:             foundUser.ID,
		Email:          foundUser.Email,
		Username:       foundUser.Username,
		FriendRequests: foundUser.FriendRequests,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userData)
}
