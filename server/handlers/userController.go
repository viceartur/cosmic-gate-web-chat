package handlers

import (
	"context"
	"cosmic-gate-chat/config"
	"cosmic-gate-chat/models"
	"cosmic-gate-chat/services"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Add New User to the Database
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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
		Username:       userRequest.Username,
		Email:          userRequest.Email,
		Password:       hashPassword,
		FriendRequests: []primitive.ObjectID{},
		Friends:        []primitive.ObjectID{},
		CreatedAt:      primitive.NewDateTimeFromTime(time.Now()),
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
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	userId := r.URL.Query().Get("userId")

	var user models.User
	var err error
	if email != "" {
		user, err = services.GetUserByEmail(email)
	} else if userId != "" {
		user, err = services.GetUserById(userId)
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Get User Friends from Database by User ID
func GetUserFriendsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	user, err := services.GetUserFriends(userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Get all Users from Database
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	users, err := services.GetUsers(userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// Send a Friend Request from a User (userId) to another User (friendId)
func SendFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	var friendRequest models.FriendRequest
	json.NewDecoder(r.Body).Decode(&friendRequest)

	err := services.SendFriendRequest(friendRequest.UserID, friendRequest.FriendID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending a friend request: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Friend request sent")
}

// Get all User Friend Requests
func GetUserFriendRequestsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	user, err := services.GetUserFriendRequests(userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Accept Friend Request from a User
func AcceptFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	var friendRequest models.FriendRequest
	json.NewDecoder(r.Body).Decode(&friendRequest)

	err := services.AcceptFriendRequest(friendRequest.UserID, friendRequest.FriendID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error adding user to friends: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Friend Request accepted")
}

// Decline Friend Request from a User
func DeclineFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	var friendRequest models.FriendRequest
	json.NewDecoder(r.Body).Decode(&friendRequest)

	err := services.DeclineFriendRequest(friendRequest.UserID, friendRequest.FriendID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error declining a friend request: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Friend Request declined")
}
