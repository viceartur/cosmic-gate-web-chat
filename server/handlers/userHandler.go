package handlers

import (
	"cosmic-gate-chat/models"
	"cosmic-gate-chat/repositories"
	"cosmic-gate-chat/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Add New User to the Database
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest models.User
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := repositories.CreateUser(userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		user, err = repositories.GetUserByEmail(email)
	} else if userId != "" {
		user, err = repositories.GetUserById(userId)
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Update User data in the Database
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData models.User
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Invalid request: %v", err),
		})
		return
	}

	if err := repositories.UpdateUser(userData); err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Error updating user: %v", err),
		})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "User updated successfully",
	})
}

// Get User Friends from Database by User ID
func GetUserFriendsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	user, err := repositories.GetUserFriends(userId)
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
	users, err := repositories.GetUsers(userId)
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

	err := repositories.SendFriendRequest(friendRequest.UserID, friendRequest.FriendID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending a friend request: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Friend request sent")
}

// Get all User Friend Requests
func GetUserFriendRequestsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	user, err := repositories.GetUserFriendRequests(userId)
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

	err := repositories.AcceptFriendRequest(friendRequest.UserID, friendRequest.FriendID)
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

	err := repositories.DeclineFriendRequest(friendRequest.UserID, friendRequest.FriendID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error declining a friend request: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Friend Request declined")
}
