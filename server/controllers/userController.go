package controllers

import (
	"context"
	"cosmic-gate-chat/config"
	"cosmic-gate-chat/models"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
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

	user := &models.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
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

// Get User from Database by UserName
func GetUser(w http.ResponseWriter, r *http.Request) {
	client := config.GetMongoDBClient()
	userCollection := client.Database("cosmic-gate-db").Collection("users")

	username := r.URL.Query().Get("username")

	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching user: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
