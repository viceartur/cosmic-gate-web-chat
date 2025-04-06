package controllers

import (
	"context"
	"cosmic-gate-chat/config"
	"cosmic-gate-chat/models"
	"cosmic-gate-chat/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Add New User to the Database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	client := config.GetMongoDBClient()
	userCollection := client.Database("cosmic-gate-db").Collection("users")

	var userRequest models.User
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashPassword, err := HashPassword(userRequest.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error hashing a password: %v", err), http.StatusInternalServerError)
		return
	}

	user := &models.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: hashPassword,
	}

	// Insert into MongoDB
	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// Get User from Database by Email
func GetUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	user, err := services.GetUserFromDB(email)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Get User Friends from Database by User ID
func GetUserFriends(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	user, err := services.GetUserFriendsFromDB(userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Get all Users from Database
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	users, err := services.GetUsersFromDB(userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
